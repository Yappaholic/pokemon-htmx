package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Type struct {
	Name string
}
type Result struct {
	Count   string
	Results []Type
}

func ApiCall() []Type {
	response, err := http.Get("http://pokeapi.co/api/v2/type")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Result
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Results
}

