package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PokemonData struct {
	Name      string   `json:"name"`
	Abilities []string `json:"abilities"`
	Weight    int      `json:"weight"`
	Types     []string `json:"types"`
}

func FetchPokemon(name string) (PokemonData, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil || resp.StatusCode != http.StatusOK {
		return PokemonData{}, errors.New("Pokemon not found")
	}
	defer resp.Body.Close()

	var data struct {
		Name      string `json:"name"`
		Abilities []struct {
			Ability struct {
				Name string `json:"name"`
			} `json:"ability"`
		} `json:"abilities"`
		Weight int `json:"weight"`
		Types  []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return PokemonData{}, errors.New("Failed to parse response")
	}

	abilities := []string{}
	for _, a := range data.Abilities {
		abilities = append(abilities, a.Ability.Name)
	}

	types := []string{}
	for _, t := range data.Types {
		types = append(types, t.Type.Name)
	}

	return PokemonData{
		Name:      data.Name,
		Abilities: abilities,
		Weight:    data.Weight,
		Types:     types,
	}, nil
}

func FetchPokemonByID(id int) (PokemonData, error) {
	return FetchPokemon(fmt.Sprintf("%d", id))
}
