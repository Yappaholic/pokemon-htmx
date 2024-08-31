package main

import (
	"htmx/api"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
type Data[T api.Result] struct {
  Count int
  ApiResult T
}
func Case (s string) string {
  caser := cases.Title(language.English)
  return caser.String(s)
}

func Id (s string) string {
  r := regexp.MustCompile("[0-9]+$")
  id := r.FindString(strings.TrimSuffix(s,"/"))
  result := "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/" + id + ".png"
  return result
}

var typesData = api.GetApiResults[api.TypeResults]("type")

var types = Data[api.TypeResults] {Count: 0, ApiResult: typesData}
