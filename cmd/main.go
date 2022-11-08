package main

import (
	"fmt"
	"github.com/perbu/asciigraph"
	"github.com/perbu/powercost/powercost"
	"os"
	"time"
)

func plot(when time.Time) error {
	zone := "NO1"
	pc, err := powercost.GetPrices(when, zone)
	if err != nil {
		return err
	}
	prices := make([]float64, len(pc))
	for i, p := range pc {
		prices[i] = p.NOKPerKWh
	}
	// find the max and min prices:
	min, max := prices[0], prices[0]
	for _, p := range prices {
		if p > max {
			max = p
		}
		if p < min {
			min = p
		}
	}
	delta := max - min
	highPrice := min + delta*0.8
	lowPrice := min + delta*0.2
	caption := fmt.Sprintf("Prices for %s in %s", when.Format("2006-01-02"), zone)
	graphs := asciigraph.Plot(prices, asciigraph.Height(10), asciigraph.Width(24*3),
		asciigraph.Caption(caption),
		asciigraph.ColorAbove(asciigraph.Red, highPrice),
		asciigraph.ColorBelow(asciigraph.DarkGreen, lowPrice))
	fmt.Println(graphs)
	fmt.Print("       ")
	for i := 0; i < 24; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%02d:00 ", i)
	}
	fmt.Println()
	fmt.Print("    ")
	for i := 0; i < 24; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Printf("%02d:00 ", i)
	}

	return nil
}

func realMain() error {
	// check if the argument tomorrow was given:
	when := time.Now()
	if len(os.Args) > 1 && os.Args[1] == "tomorrow" {
		when = when.Add(24 * time.Hour)
	}
	err := plot(when)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := realMain()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
