package route

import (
	"github.com/karthikeswar17/buzz/message/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func MessageRoute(e *echo.Echo, DB *mongo.Client) {
	var messageHandler handler.MessageHandler = handler.MessageHandler{DB: *DB}
	e.POST("message", messageHandler.AddMessage)
	e.GET("conversation/:id/message/list", messageHandler.GetMessages)
}
