package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	render "gitlab.com/lenivyyluitel/gahara/pkg/renderer"
	"gitlab.com/lenivyyluitel/gahara/routes"
)

func main() {
	app := echo.New()
	app.HideBanner = true

	t := &render.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("templates/*.*")),
	}
	app.Renderer = t

	app.GET("/", routes.Home)
	app.GET("/image", routes.RandomFiles)

	err := app.Start(":1234")
	if err != nil {
		app.Logger.Fatal(err)
	}
}
