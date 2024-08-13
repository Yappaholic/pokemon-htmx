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
type TypeResults struct {
  Results []Type
}

type PokemonList struct {
  Pokemon struct{Name string}
}

type TypePokemonResult struct {
  Pokemon []PokemonList
}

type Result interface {
  TypeResults | TypePokemonResult
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
