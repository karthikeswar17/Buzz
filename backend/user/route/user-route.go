package route

import (
	"github.com/karthikeswar17/user/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(e *echo.Echo, DB *mongo.Client) {
	var handler handler.UserHandler
	handler.DB = *DB
	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
}
