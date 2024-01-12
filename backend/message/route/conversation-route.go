package route

import (
	"github.com/karthikeswar17/buzz/message/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConversationRoute(e *echo.Echo, DB *mongo.Client) {
	var conversationHandler handler.ConversationHandler = handler.ConversationHandler{DB: *DB}
	e.GET("conversation/:id/message/list", conversationHandler.GetMessages)
	e.POST("conversation", conversationHandler.CreateConversation)
}
