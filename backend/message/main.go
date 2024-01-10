package main

import (
	"net/http"

	"github.com/karthikeswar17/buzz/message/route"
	"github.com/karthikeswar17/buzz/message/util"
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
		SigningKey:  []byte(util.SigningKey),
		TokenLookup: "header:Authorization,cookie:token",
	}))

	DB := util.ConnectDB()
	route.MessageRoute(e, DB)
	e.Logger.Fatal(e.Start(":8002"))
}
