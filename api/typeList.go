package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)
type PokemonInfo struct {
  Name string
  Url string
}
type Pokemon struct {
  Pokemon struct{Name string; Url string}
}

type PokemonList struct {
  Pokemon []Pokemon 
}

func PokemonCall(route string) []Pokemon {
	response, err := http.Get("http://pokeapi.co/api/v2/type/" + route)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject PokemonList
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Pokemon
}



