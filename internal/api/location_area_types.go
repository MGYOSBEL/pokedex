package api

type LocationAreaListResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedApiResource `json:"results"`
}

type LocationAreaResponse struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod NamedApiResource `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int              `json:"rate"`
			Version NamedApiResource `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location NamedApiResource `json:"location"`
	Names    []struct {
		Name     string           `json:"name"`
		Language NamedApiResource `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        NamedApiResource `json:"pokemon"`
		VersionDetails []struct {
			Version          NamedApiResource `json:"version"`
			MaxChance        int              `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int              `json:"min_level"`
				MaxLevel        int              `json:"max_level"`
				ConditionValues []any            `json:"condition_values"`
				Chance          int              `json:"chance"`
				Method          NamedApiResource `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
