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

func plot(when time.Time, zone string, ore, nomva, showCurrentHour bool) error {
	pc, err := powercost.GetPrices(when, zone, nomva)
	if err != nil {
		return fmt.Errorf("powercost.GetPrices: %w", err)
	}
	prices := make([]float64, len(pc))
	for i, p := range pc {
		if ore {
			prices[i] = p.NOKPerKWh * 100
		} else {
			prices[i] = p.NOKPerKWh
		}
	}
	// find the maxPrice and minPrice prices:
	minPrice, maxPrice := prices[0], prices[0]
	for _, p := range prices {
		if p > maxPrice {
			maxPrice = p
		}
		if p < minPrice {
			minPrice = p
		}
	}
	delta := maxPrice - minPrice
	highPrice := minPrice + delta*0.8
	lowPrice := minPrice + delta*0.2
	unit := "Kr/kWh"
	if ore {
		unit = "øre/kWh"
	}
	if nomva {
		unit = unit + " excl MVA"
	}
	caption := fmt.Sprintf("Prices for %s in %s (%s)", when.Format("2006-01-02"), zone, unit)
	graphs := asciigraph.Plot(prices, asciigraph.Height(10), asciigraph.Width(24*3),
		asciigraph.Caption(caption),
		asciigraph.ColorAbove(asciigraph.Red, highPrice),
		asciigraph.ColorBelow(asciigraph.DarkGreen, lowPrice))
	fmt.Println(graphs)
	// find the first vertical line so we can align the x-axis labels
	margin := strings.IndexAny(graphs, "┤┼") - 1
	printXaxisLabels(margin, showCurrentHour)
	return nil
}

func printXaxisLabels(margin int, showCurrentHour bool) {
	fmt.Print(strings.Repeat(" ", margin))
	// extract the current hour
	hour := time.Now().Hour()
	for i := 0; i <= 24; i++ {
		if showCurrentHour && i == hour {
			// print the current hour in reverse video
			fmt.Printf("\033[7m%02d\033[0m ", i)
		} else {
			fmt.Printf("%02d ", i)
		}
	}
	fmt.Println()
}

func realMain() error {
	tomorrow := flag.Bool("tomorrow", false, "Show price for tomorrow instead of today")
	yesterday := flag.Bool("yesterday", false, "Show price for yesterday instead of today")
	ore := flag.Bool("ore", false, "Show price in øre, instead of NOK")
	date := flag.String("date", "", "Show price for the given date (YYYY-MM-DD)")
	zone := flag.String("zone", "NO1", "Which price zone to show")
	nomva := flag.Bool("no-mva", false,
		"Always from prices without MVA. Default is to include it for NO1, NO2 and NO3. NO4 pays no MVA.")
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
			return fmt.Errorf("invalid date: %w", err)
		}
	}
	var showCurrentHour bool
	if !*tomorrow || !*yesterday || *date == "" {
		showCurrentHour = true
	}

	// make sure *zone is uppercase:
	*zone = strings.ToUpper(*zone)
	if *zone == "NO4" { // no MVA in NO4
		*nomva = true
	}
	err := plot(when, *zone, *ore, *nomva, showCurrentHour)
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
