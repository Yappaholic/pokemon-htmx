package main

import (
	"htmx/api"
	"htmx/globals"
	"htmx/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func homeHandler(c echo.Context) error {
	return render(c, views.Index())
}
func countHandler(c echo.Context) error {
	globals.Number++
	return render(c, views.Count(globals.Number))
}

func apiHandler(c echo.Context) error {
	return render(c, views.ApiResult(api.ApiCall()))
}
func render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}
func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.Use(middleware.Logger())
	e.GET("/", homeHandler)
	e.POST("/count", countHandler)
	e.Logger.Fatal(e.Start(":6969"))
}
