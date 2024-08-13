package main

import (
  "htmx/api"
)
type Data[T api.Result] struct {
  Count int
  ApiResult T
}
var typesData = api.GetApiResults[api.TypeResults]("type")

var types = Data[api.TypeResults] {Count: 0, ApiResult: typesData}
