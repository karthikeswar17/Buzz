package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/karthikeswar17/buzz/message/model"
	"github.com/karthikeswar17/buzz/message/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageHandler struct {
	DB mongo.Client
}

func (handler MessageHandler) AddMessage(c echo.Context) error {
	var err error
	var message model.Message
	err = (&echo.DefaultBinder{}).BindBody(c, &message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed"})
	}
	message.Id = uuid.NewString()
	messageCollection := util.GetCollection(&handler.DB, "Message")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = messageCollection.InsertOne(ctx, message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "Success"})
}

func (handler MessageHandler) GetMessages(c echo.Context) error {
	var conversationId string = c.Param("id")
	var messages []model.Message
	messageCollection := util.GetCollection(&handler.DB, "Message")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := messageCollection.Find(ctx, bson.M{"conversation": conversationId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	cursor.All(ctx, &messages)
	return c.JSON(http.StatusOK, messages)

}
