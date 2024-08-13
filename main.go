package main

import (
	"html/template"
	"htmx/api"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type Templates struct {
  templates *template.Template
}
func (t *Templates) Render (w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data) 
}

func NewTemplate () *Templates {
  funcMap := template.FuncMap {
    "Case": Case,
  }
  return &Templates {
    templates: template.Must(template.New("templ").Funcs(funcMap).ParseGlob("views/*.html")) ,
  }
}
func homeHandler(c echo.Context) error {
	return c.Render(200, "index", types)
}
func typeHandler(c echo.Context) error {
	//instantiating struct for template rendering
	type d struct {List []api.PokemonList; Case func(string) string}
	typeList := api.GetApiResults[api.TypePokemonResult]("type/", c)
	return c.Render(200, "typeList", d{List: typeList.Pokemon, Case: Case})
}

func pokemonHandler(c echo.Context) error {
	type d struct {Pokemon api.Pokemon; Case func(string) string}
	typeList := api.GetApiResults[api.Pokemon]("pokemon/", c)
	return c.Render(200, "pokemon", d{Pokemon: typeList, Case: Case})
}

func main() {
  e := echo.New();
  e.Use(middleware.Logger())
  e.Static("/static", "static") 
  e.Renderer = NewTemplate()
  e.GET("/", homeHandler)
  e.GET("/type/:name", typeHandler)
  e.GET("/pokemon/:name", pokemonHandler)
  e.Logger.Fatal(e.Start(":6969"))
  
}
