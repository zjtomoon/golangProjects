package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	render := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = render
	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "something.html", map[string]interface{}{
			"name": "Dolly",
		})
	}).Name = "foobar"

	//log.Fatal(e.Start(":3000"))
	e.Logger.Fatal(e.Start(":3000"))
}
