package api

import (
  "encoding/json"
  "net/http"
  "io"
)
type Type struct {
  Name string
}
type TypeResults struct {
  Results []Type
}
func  GetTypes () []Type  {
  response, err := http.Get("http://pokeapi.co/api/v2/type")
  if err != nil {
    return nil
  }
  responseData, err := io.ReadAll(response.Body)
  if err != nil {
    return nil
  }
  var result TypeResults
  json.Unmarshal(responseData, &result)
  return result.Results
}
