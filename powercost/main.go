package powercost

import (
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"time"
)

type PowerCost struct {
	NOKPerKWh float64   `json:"NOK_per_kWh"`
	EURPerKWh float64   `json:"EUR_per_kWh"`
	Exr       float64   `json:"EXR"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
}

// GetPrices for the given date.
func GetPrices(when time.Time, zone string) ([]PowerCost, error) {
	var powerCost []PowerCost
	err := requests.
		URL(getUrl(when, zone)).
		ToJSON(&powerCost).
		Fetch(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("fetching power cost: %w", err)
	}
	return powerCost, nil
}

// GetCurrent takes a list of PowerCost and returns the current price
func GetCurrent(pc []PowerCost) (PowerCost, error) {
	// iterate over the list and return the first one that is in the future
	for _, p := range pc {
		if p.TimeStart.After(time.Now()) {
			return p, nil
		}
	}
	return PowerCost{}, fmt.Errorf("no current price found")
}

func getUrl(date time.Time, zone string) string {
	// https://www.hvakosterstrommen.no/api/v1/prices/2022/11-08_NO5.json
	return fmt.Sprintf("https://www.hvakosterstrommen.no/api/v1/prices/%d/%s_%s.json", date.Year(), date.Format("01-02"), zone)
}
