package router

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	
	"github.com/yamato0204/go-up/app/usecase"
)


type Template struct {
	templates *template.Template
  }

  func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func GetRouter(u usecase.Usecase) *echo.Echo {

	e := echo.New()

	renderer := &Template{
        templates: template.Must(template.ParseGlob("view/*.html")),
    }

	e.Renderer = renderer

	e.GET("/", usecase.GetTop)
	e.GET("/signup", usecase.GetSignup)
	
	e.GET("/login", usecase.GetLogin)
	e.POST("/signup", u.PostSignup)
	e.POST("login", u.PostLogin )


	return e
}