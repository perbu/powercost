package main

import (
	"flag"
	"fmt"
	"github.com/perbu/asciigraph"
	"github.com/perbu/powercost/powercost"
	"os"
	"strings"
	"time"
)

func plot(when time.Time, zone string, ore bool, mva bool) error {
	pc, err := powercost.GetPrices(when, zone, mva)
	if err != nil {
		return err
	}
	prices := make([]float64, len(pc))
	for i, p := range pc {
		if ore {
			prices[i] = p.NOKPerKWh * 100
		} else {
			prices[i] = p.NOKPerKWh
		}
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
	unit := "Kr/kWh"
	if ore {
		unit = "øre/kWh"
	}
	caption := fmt.Sprintf("Prices for %s in %s (%s)", when.Format("2006-01-02"), zone, unit)
	graphs := asciigraph.Plot(prices, asciigraph.Height(10), asciigraph.Width(24*3),
		asciigraph.Caption(caption),
		asciigraph.ColorAbove(asciigraph.Red, highPrice),
		asciigraph.ColorBelow(asciigraph.DarkGreen, lowPrice))
	fmt.Println(graphs)
	margin := strings.Index(graphs, "┤") - 1
	printXaxisLabels(margin)
	return nil
}

func printXaxisLabels(margin int) {
	fmt.Print(strings.Repeat(" ", margin))

	for i := 0; i <= 24; i++ {
		fmt.Printf("%02d ", i)
	}
	fmt.Println()
}

func realMain() error {
	tomorrow := flag.Bool("tomorrow", false, "Show price for tomorrow instead of today")
	yesterday := flag.Bool("yesterday", false, "Show price for yesterday instead of today")
	ore := flag.Bool("ore", false, "Show price in øre, instead of NOK")
	date := flag.String("date", "", "Show price for the given date (YYYY-MM-DD)")
	zone := flag.String("zone", "NO1", "Which price zone to show")
	mva := flag.Bool("mva", false, "Show prices including MVA (25%)")
	flag.Parse()
	if *tomorrow && *yesterday {
		return fmt.Errorf("Can't show both yesterday and tomorrow")
	}
	// check if the argument tomorrow was given:
	when := time.Now()
	if *tomorrow {
		when = when.Add(24 * time.Hour)
	}
	if *yesterday {
		when = when.Add(-24 * time.Hour)
	}
	if *date != "" {
		if *tomorrow || *yesterday {
			return fmt.Errorf("Can't show both a specific date and yesterday or tomorrow")
		}
		var err error
		when, err = time.Parse("2006-01-02", *date)
		if err != nil {
			return fmt.Errorf("Invalid date: %v", err)
		}
	}

	// make sure *zone is uppercase:
	*zone = strings.ToUpper(*zone)
	err := plot(when, *zone, *ore, *mva)
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
