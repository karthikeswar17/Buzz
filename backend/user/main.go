package main

import (
	"net/http"

	"github.com/karthikeswar17/user/route"
	"github.com/karthikeswar17/user/util"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(util.SigningKey),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
		TokenLookup: "header:Authorization,cookie:token",
	}))

	DB := util.ConnectDB()
	route.UserRoute(e, DB)
	route.FriendRoute(e, DB)
	e.Logger.Fatal(e.Start(":8000"))
}
