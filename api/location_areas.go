package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const LOCATION_AREA_URL = API_URL + "/location-area/"

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("response failed with status code %d: %s", res.StatusCode, body)
	}
	if err != nil {
		return LocationArea{}, err
	}
	var la LocationArea
	json.Unmarshal(body, &la)
	return la, nil
}
