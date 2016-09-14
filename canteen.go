package openmensa

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// A Canteen object
type Canteen struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	Address    string    `json:"address"`
	Coordinate []float64 `json:"coordinates"`
}

func (c Canteen) String() string {
	return fmt.Sprintf("%s", c.Name)
}

// GetAllCanteens returns a list of all known canteens
// Since OpenMensa uses pagination for the results, this will result in multiple requests.
func GetAllCanteens() (canteens []*Canteen, err error) {
	// FIXME: Concurrency anyone?

	url, _ := url.Parse(BaseURL + "canteens")
	page := 1

	for {
		params := url.Query()
		params.Set("page", fmt.Sprintf("%d", page))
		url.RawQuery = params.Encode()

		resp, err := get(url.String())

		var currentCanteens []*Canteen
		err = json.Unmarshal(resp, &currentCanteens)
		if err != nil {
			return nil, err
		}

		if len(currentCanteens) == 0 {
			// Reached the last page
			break
		}

		canteens = append(canteens, currentCanteens...)

		page++
	}

	return
}

// GetCanteen returns a canteen object for the given id
func GetCanteen(id int) (canteen *Canteen, err error) {
	url := fmt.Sprintf("%s/canteens/%d", BaseURL, id)
	resp, err := get(url)

	err = json.Unmarshal(resp, &canteen)

	return
}
