package usecase

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/go-up/app/entity"
	"github.com/yamato0204/go-up/app/infra"
)


type Usecase interface {
	
	PostSignup(c echo.Context) error
	PostLogin(c echo.Context)error
	GetHome(c echo.Context) error
	Auth(c echo.Context, userID string) (entity.User, error)
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

func (u *usecase)GetHome(c echo.Context) error {
    user := entity.User{}
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	
	 cookieKey := os.Getenv("LOGIN_USER_ID_KEY")

	 userId, err := infra.GetSession(c, cookieKey)
	
	user, err =  u.Auth(c,userId)
	if err != nil {
		fmt.Println("auth error 2")
		return err
	}

	fmt.Println(user)

	// err =  c.Render(http.StatusOK, "home", map[string]interface{}{"user": user}); 
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	return nil
	
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
	c.Redirect(http.StatusFound, "/home")
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

	// fmt.Println(pw)
	// fmt.Println(pass)
	if pw != pass {
		c.Redirect(http.StatusFound, "/login")
	}
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	infra.NewSession(c, cookieKey, storeUser.ID)
	c.Redirect(http.StatusFound, "/home")
	return nil
}

func (u *usecase)Auth(c echo.Context, userID string ) (entity.User, error) {
	 user := entity.User{}
	 err := u.i.GetOneUser(&user, userID);

	 if err != nil {
		fmt.Println("auth error")
	 c.Redirect(301, "/login")
	 return user, err
		
	}

	
	return user, nil
		
	
		
	}

	




