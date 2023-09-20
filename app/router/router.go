package router

import (
	"html/template"
	"io"
	"net/http"

	"os"

	"github.com/labstack/echo/v4"

	"github.com/yamato0204/go-up/app/infra"
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

	// loginCheck := e.Group("/")
	// loginCheck.Use(checkLogin())

	// loginCheck.GET("/h", u.GetHome)

    e.GET("/home", u.GetHome)
	e.GET("/top", usecase.GetTop)
	e.GET("/signup", usecase.GetSignup)
	
	e.GET("/login", usecase.GetLogin)
	e.POST("/signup", u.PostSignup)
	e.POST("login", u.PostLogin )
	e.GET("/logout", u.Logout)


	return e
}

func checkLogin()echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	id, err := infra.GetSession(c, cookieKey)
	if err != nil {
		return err
	}
	if id == "" {
			c.Redirect(http.StatusFound, "/login")
		} else {
			//c.NoContent(http.StatusOK)
		}
		return next(c)
    }
}
	
}