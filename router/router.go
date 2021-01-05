package router

import (
	"errors"
	"net/http"

	"github.com/higashi000/spovimfy/oauth2config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/oauth2"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler)
	e.GET("/callback/", callbackHandler)

	return e
}

func handler(c echo.Context) error {
	url := oauth2config.Config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	return c.Redirect(301, url)
}

func callbackHandler(c echo.Context) error {
	token, err := oauth2config.Config.Exchange(oauth2.NoContext, c.Request().URL.Query().Get("code"))
	if err != nil {
		return err
	}

	if !token.Valid() {
		return errors.New("vaild token")
	}

	return c.String(http.StatusOK, token.AccessToken)
}
