package route

import (
	"github.com/karthikeswar17/buzz/socket-server/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func SocketRoute(e *echo.Echo, DB *mongo.Client) {
	var handler handler.SocketHandler
	handler.DB = *DB
	e.GET("/ws", handler.WsHandler)

}
