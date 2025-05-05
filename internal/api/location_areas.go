package api

import (
	"encoding/json"
	"fmt"
)

const LOCATION_AREA_URL = API_URL + "/location-area/"

func (c *Client) GetLocationAreas(url *string) (LocationAreasResponse, error) {
	endpoint := LOCATION_AREA_URL
	if url != nil {
		endpoint = *url
	}
	var la LocationAreasResponse

	response, err := c.GetCached(endpoint)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	err = json.Unmarshal(response, &la)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return la, nil
}

func (c *Client) GetLocationAreaByName(name *string) (LocationAreaResponse, error) {
	if name == nil {
		return LocationAreaResponse{}, fmt.Errorf("name can not be empty")
	}

	endpoint := fmt.Sprintf("%s/%s", LOCATION_AREA_URL, *name)
	var la LocationAreaResponse

	result, err := c.GetCached(endpoint)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	err = json.Unmarshal(result, &la)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return la, nil
}
