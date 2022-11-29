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
func GetPrices(when time.Time, zone string, mva bool) ([]PowerCost, error) {
	var powerCost []PowerCost
	err := requests.
		URL(getUrl(when, zone)).
		ToJSON(&powerCost).
		Fetch(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("fetching power cost: %w", err)
	}
	// if we want prices including MVA we adjust it here.
	if mva {
		for i, p := range powerCost {
			powerCost[i].NOKPerKWh = p.NOKPerKWh * 1.25
		}
	}
	return powerCost, nil
}

func getUrl(date time.Time, zone string) string {
	// https://www.hvakosterstrommen.no/api/v1/prices/2022/11-08_NO5.json
	return fmt.Sprintf("https://www.hvakosterstrommen.no/api/v1/prices/%d/%s_%s.json", date.Year(), date.Format("01-02"), zone)
}
