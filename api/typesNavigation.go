package api

import (
  "encoding/json"
  "net/http"
  "github.com/labstack/echo/v4"
  "io"
  "os"
)

type Type struct {
  Name string
}
//Results in types list for sidebar navigation
type TypeResults struct {
  Results []Type
}

type PokemonList struct {
  Pokemon struct{Name string; Url string}
}
//Results in pokemon list of specified type 
type TypePokemonResult struct {
  Pokemon []PokemonList
}
//Results in stats for a specific pokemon
type TypeStats struct {
  Base_stat int
  Stat struct {Name string}
}
type Pokemon struct {
  Name string
  Stats []TypeStats
  Sprites struct{Front_default string}
}

type Result interface {
  TypeResults | TypePokemonResult | Pokemon
}
func GetApiResults [T Result] (query string, c... echo.Context) T {
  var url string
  if len(c) != 0 {
  	url = "http://pokeapi.co/api/v2/" + query + c[0].Param("name")
  } else {
    url = "http://pokeapi.co/api/v2/" + query 
  }
  response, err := http.Get(url)
  if err != nil {
    os.Exit(1) 
  }
  responseData, err := io.ReadAll(response.Body)
  var result T
  json.Unmarshal(responseData, &result)
  return result
}
