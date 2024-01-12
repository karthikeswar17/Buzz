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

type ConversationHandler struct {
	DB mongo.Client
}

func (handler ConversationHandler) GetMessages(c echo.Context) error {
	var conversationId string = c.Param("id")
	var messages []model.Message = []model.Message{}
	messageCollection := util.GetCollection(&handler.DB, "Message")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := messageCollection.Find(ctx, bson.M{"conversation_id": conversationId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	cursor.All(ctx, &messages)
	return c.JSON(http.StatusOK, messages)

}
func (handler *ConversationHandler) CreateConversation(c echo.Context) error {
	var err error
	var conversationRequest model.ConversationRequest
	err = (&echo.DefaultBinder{}).BindBody(c, &conversationRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	var conversation model.Conversation
	var userId string = util.GetUserIdFromJWT(c)
	var conversationCollection = util.GetCollection(&handler.DB, "Conversation")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var members []string = []string{userId}
	members = append(members, conversationRequest.Members...)
	err = conversationCollection.FindOne(ctx, bson.M{"members": bson.M{"$all": members}}).Decode(&conversation)
	if err != nil && err == mongo.ErrNoDocuments {
		conversation = model.Conversation{
			Id:      uuid.NewString(),
			Members: members,
		}
		_, err = conversationCollection.InsertOne(ctx, conversation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
		}
	}
	return c.JSON(http.StatusOK, conversation)
}
