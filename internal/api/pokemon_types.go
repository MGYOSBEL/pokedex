package api

import "fmt"

type Pokemon struct {
	Id             int              `json:"id"`
	Name           string           `json:"name"`
	BaseExperience int              `json:"base_experience"`
	Height         int              `json:"height"`
	IsDefault      bool             `json:"is_default"`
	Weight         int              `json:"weight"`
	Order          int              `json:"order"`
	Abilities      []PokemonAbility `json:"abilities"`
	Stats          []PokemonStat    `json:"stats"`
	Types          []PokemonType    `json:"types"`
}

type PokemonAbility struct {
	Ability  NamedApiResource `json:"ability"`
	IsHidden bool             `json:"is_hidden"`
	Slot     int              `json:"slot"`
}

type PokemonStat struct {
	Stat     NamedApiResource `json:"stat"`
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
}

type PokemonType struct {
	Type NamedApiResource `json:"type"`
	Slot int              `json:"slot"`
}

func (p Pokemon) String() string {
	ps := fmt.Sprintf("Name: %s\n", p.Name)
	ps = fmt.Sprintf("%sHeight: %d\n", ps, p.Height)
	ps = fmt.Sprintf("%sWeight: %d\n", ps, p.Weight)
	ps = fmt.Sprintf("%sStats:\n", ps)
	for _, stat := range p.Stats {
		ps = fmt.Sprintf("%s  -%s: %d\n", ps, stat.Stat.Name, stat.BaseStat)
	}

	ps = fmt.Sprintf("%sTypes:\n", ps)
	for _, ty := range p.Types {
		ps = fmt.Sprintf("%s  - %s\n", ps, ty.Type.Name)
	}
	return ps
}
