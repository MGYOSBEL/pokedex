package api

import (
	"encoding/json"
	"fmt"
)

const POKEMON_URL = API_URL + "/pokemon"

func (c *Client) GetPokemonByName(name *string) (Pokemon, error) {
	if name == nil {
		return Pokemon{}, fmt.Errorf("name can not be empty")
	}

	endpoint := fmt.Sprintf("%s/%s", POKEMON_URL, *name)
	var pokemon Pokemon

	res, err := c.GetCached(endpoint)
	if err != nil {
		return Pokemon{}, err
	}
	err = json.Unmarshal(res, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
