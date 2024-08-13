package main

import (
  "htmx/api"
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

var typesData = api.GetApiResults[api.TypeResults]("type")

var types = Data[api.TypeResults] {Count: 0, ApiResult: typesData}
