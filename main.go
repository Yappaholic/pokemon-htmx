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
  return &Templates {
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}
func homeHandler(c echo.Context) error {
	return c.Render(200, "index", types)
}
func typeHandler(c echo.Context) error {
	type d struct {List []api.PokemonList}
	typeList := api.GetApiResults[api.TypePokemonResult]("type/", c)
	return c.Render(200, "typeList", d{List: typeList.Pokemon})
}

func main() {
  e := echo.New();
  e.Use(middleware.Logger())
  e.Static("/static", "static") 
  e.Renderer = NewTemplate()
  e.GET("/", homeHandler)
  e.GET("/type/:name", typeHandler)
  e.Logger.Fatal(e.Start(":6969"))
  
}
