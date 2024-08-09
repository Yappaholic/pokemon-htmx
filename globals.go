package main

import (
  "htmx/api"
)
type Data struct {
  Count int
  Types []api.Type
}
var data = Data {Count: 0, Types: api.GetTypes()}
