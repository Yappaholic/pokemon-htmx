package api
import (
  "encoding/json"
  "net/http"
  "io"
  "github.com/labstack/echo/v4"
)
type PokemonList struct {
  Pokemon struct{Name string}
}

type TypePokemonResult struct {
  Pokemon []PokemonList
}
func  GetTypePokemons (c echo.Context) []PokemonList  {
  query := c.Param("name") 
  response, err := http.Get("http://pokeapi.co/api/v2/type/" + query)
  if err != nil {
    return nil
  }
  responseData, err := io.ReadAll(response.Body)
  if err != nil {
    return nil
  }
  var result TypePokemonResult
  json.Unmarshal(responseData, &result)
  return result.Pokemon
}


