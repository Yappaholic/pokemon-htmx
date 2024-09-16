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
    "Id": Id,
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
	type d struct {List []api.PokemonList; Case func(string) string; Type string}
	typeList := api.GetApiResults[api.TypePokemonResult]("type/", c)
	return c.Render(200, "typeList", d{List: typeList.Pokemon, Case: Case, Type: c.Param("name")})
}

func pokemonHandler (c echo.Context) error  {
	type d struct {Pokemon api.Pokemon; Case func(string) string}
	type e struct {Error string}
	typeList := api.GetApiResults[api.Pokemon]("pokemon/", c)
  if c.Request().Header["HX-Request"] != nil {
    if typeList.Name != "" {
    	return c.Render(200, "pokemon", d{Pokemon: typeList, Case: Case})
    } else {
    	return c.Render(303, "error", e{Error: "Pokemon not Found!"})
    }
  } else {
    return pokemonSelfHandler(c, typeList)
  }
}  

func pokemonSelfHandler (c echo.Context, typelist api.Pokemon) (error,error) {
  	type d struct {Pokemon api.Pokemon; Case func(string) string}
  	type e struct {Error string}
  	if typelist.Name != "" {
    	return c.Render(200, "index", types), c.Render(200, "pokemon", d{Pokemon: typelist, Case: Case})
  	} else {
  	  c.Render(200, "index", types)
    	return c.Render(200,"index", types), c.Render(303, "error", e{Error: "Pokemon not Found!"})
  	}
  
}
func searchHandler(c echo.Context) error {
  search := "/pokemon/" + c.QueryParam("name")
  return c.Redirect(302,search)
}

func main() {
  e := echo.New();
  e.Use(middleware.Logger())
  e.Static("/static", "static") 
  e.Renderer = NewTemplate()
  e.GET("/", homeHandler)
  e.GET("/type/:name", typeHandler)
  e.GET("/pokemon/:name", pokemonHandler)
  e.GET("/search", searchHandler)
  e.Logger.Fatal(e.Start(":6969"))
  
}
