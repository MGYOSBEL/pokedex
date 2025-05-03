package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const LOCATION_AREA_URL = API_URL + "/location-area/"

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(url *string) (LocationAreaResponse, error) {
	endpoint := LOCATION_AREA_URL
	if url != nil {
		endpoint = *url
	}
	var la LocationAreaResponse

	if cached, ok := c.cc.Get(endpoint); ok {
		err := json.Unmarshal(cached, &la)
		if err == nil {
			return la, nil
		}
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	res, err := c.Client.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("response failed with status code %d: %s", res.StatusCode, body)
	}
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cc.Add(endpoint, body)

	err = json.Unmarshal(body, &la)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return la, nil
}
