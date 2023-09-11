package usecase

import (
	
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/go-up/app/entity"
	"github.com/yamato0204/go-up/app/infra"
)


type Usecase interface {
	
	PostSignup(c echo.Context) error
	PostLogin(c echo.Context)error
}

type usecase struct {
	i infra.Infra
}

func NewUsecase(i infra.Infra) Usecase {
	return &usecase{i}
}


func GetTop(c echo.Context) error{

	topData := "Top"

	return c.Render(http.StatusOK, "top", topData )
}


func GetSignup(c echo.Context) error {

	return c.Render(http.StatusOK, "signup", nil)
}

func GetLogin(c echo.Context) error {
	
	return c.Render(http.StatusOK, "login", nil)
}

func (u *usecase)PostSignup(c echo.Context) error {

	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.Email = c.Request().PostFormValue("email")
	user.Password = c.Request().PostFormValue("password")

	user.ID = uuid.New().String()
	
	err := u.i.CreateUser(&user)
	if err != nil {
		c.Redirect(301, "/signup")
		return nil
	}
	c.Redirect(http.StatusFound, "/")
	return nil
}

func (u *usecase)PostLogin(c echo.Context) error {
	email := c.Request().PostFormValue("email")
	pw := c.Request().PostFormValue("password")

	storeUser := entity.User{}
	 err := u.i.GetUserByEmail(&storeUser, email)
	if err != nil {
		return err
	}

	pass := storeUser.Password

	if pw != pass {
		c.Redirect(http.StatusFound, "/login")
	}

	c.Redirect(http.StatusFound, "/")
	return nil
}



