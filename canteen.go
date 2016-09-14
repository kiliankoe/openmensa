package openmensa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// GetCanteen returns a canteen object for the given id
func GetCanteen(id int) (canteen *Canteen, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/canteens/%d", BaseURL, id))
	if err != nil {
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(respData, &canteen)

	return
}
