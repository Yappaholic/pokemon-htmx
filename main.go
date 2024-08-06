package main

import (
	"htmx/api"
	"htmx/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"math/rand/v2"
	"strconv"
)

func homeHandler(c echo.Context) error {
	return render(c, views.Index())
}

func apiHandler(c echo.Context) error {
	return render(c, views.ApiResult(api.ApiCall()))
}
func typeHandler(c echo.Context) error {
  route := c.ParamValues()[0]
	return render(c, views.TypePage(c, api.PokemonCall(route)))
}
func searchHandler(c echo.Context) error {
  if c.FormValue("random") == "true" {
    return c.Redirect(200, "/search/lucky") 
  } 
  query := c.FormValue("query")
  searchResult := api.PokemonSearch(query) 
  return render(c, views.SearchCall(searchResult))
  
}
func luckyHandler(c echo.Context) error {
  id := strconv.Itoa(rand.IntN(100))
  searchResult := api.PokemonSearch(id) 
  return render(c, views.SearchCall(searchResult))
}
func render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}
func main() {
	e := echo.New()
	e.Static("/static", "./static")
	e.Use(middleware.Logger())
	e.GET("/", homeHandler)
	e.GET("/type/:id", typeHandler)
	e.POST("/search", searchHandler)
	e.GET("/search/lucky", luckyHandler)
	e.Logger.Fatal(e.Start(":6969"))
}
