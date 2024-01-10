package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/karthikeswar17/buzz/socket-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetConversationFromId(DB mongo.Client, ConversationId string) model.Conversation {
	conversationCollection := GetCollection(&DB, "Conversation")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var conversation model.Conversation
	conversationCollection.FindOne(ctx, bson.M{"_id": ConversationId}).Decode(&conversation)
	return conversation
}

func SendMessage(message model.Message, token *jwt.Token) error {
	fmt.Println("sending message")
	posturl := "http://message:8002/message"
	// JSON body
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// Create a HTTP post request
	request, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	claims := getClaimsFromJWT(token)
	newToken, _ := CreateJWT(claims["user-id"].(string), claims["email"].(string), claims["tag"].(string))
	request.Header.Add("Authorization", newToken)
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}
