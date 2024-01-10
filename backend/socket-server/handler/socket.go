package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/karthikeswar17/buzz/socket-server/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/internal/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientManager struct {
	clientsMapping map[string][]*Client
}

var manager = ClientManager{
	clientsMapping: make(map[string][]*Client),
}

type Client struct {
	Id      string
	Socket  *websocket.Conn
	userTag string
	Context echo.Context
}
type Message struct {
	Id             string `json:"_id" bson:"_id" validate:"required"`
	From           string `json:"from" validate:"required"`
	To             string `json:"to" validate:"required"`
	Message        string `json:"message" validate:"required"`
	ConversationId string `json:"conversation_id" bson:"conversation_id" validate:"required"`
}

func (c *Client) read() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		//Read message
		_, rawMessage, err := c.Socket.ReadMessage()
		fmt.Println(rawMessage)
		//If there is an error message, cancel this connection and then close it
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		message := Message{Id: uuid.New().String(), From: c.userTag}
		stringMessage := string(rawMessage[:])
		err = json.Unmarshal([]byte(stringMessage), &message)
		if err != nil {
			log.Println(err.Error())
		}
		//adding to DB
		DB := c.Context.Get("DB").(*mongo.Client)
		messageCollection := util.GetCollection(DB, "Message")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Println(message)
		_, err = messageCollection.InsertOne(ctx, message)
		if err != nil {
			log.Println(err.Error())
		}
		manager.send(message)

	}
}
func WsHandler(c echo.Context) error {
	var (
		upgrader = websocket.Upgrader{}
	)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	//creating in memory user-tag to client mapping
	_, userTag := util.GetEmailAndTagFromJWT(c)
	fmt.Println(userTag)
	client := &Client{Id: uuid.New(), Socket: ws, userTag: userTag, Context: c}
	manager.clientsMapping[userTag] = append(manager.clientsMapping[userTag], client)
	go client.read()

	// for {
	// 	// Write
	// 	err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
	// 	if err != nil {
	// 		c.Logger().Error(err)
	// 	}

	// 	// Read
	// 	_, msg, err := ws.ReadMessage()
	// 	if err != nil {
	// 		c.Logger().Error(err)
	// 	}
	// 	fmt.Printf("%s\n", msg)
	// }
}
