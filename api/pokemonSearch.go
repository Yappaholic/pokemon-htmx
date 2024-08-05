package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)
type Ability struct {
  AbilityName struct{Name string}
}
type Move struct {
  MoveName struct {Name string}
}
type PokemonResult struct {
  Name string;
  Abilities []Ability;
  Moves []Move
}

func PokemonSearch(query string) PokemonResult {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + query)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject PokemonResult
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}



