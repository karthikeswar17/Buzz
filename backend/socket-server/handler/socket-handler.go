package handler

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/karthikeswar17/buzz/socket-server/model"
	"github.com/karthikeswar17/buzz/socket-server/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type SocketHandler struct {
	DB mongo.Client
}

type ClientManager struct {
	clientsMapping map[string][]*Client
	DB             mongo.Client
}

var manager = ClientManager{
	clientsMapping: make(map[string][]*Client),
}

func (manager *ClientManager) send(message model.Message) {
	fmt.Println(message)
	var conversation model.Conversation = util.GetConversationFromId(manager.DB, message.ConversationId)
	// var receivers []string
	for _, member := range conversation.Members {
		if member != message.From {
			// receivers = append(receivers, member)
			toClients := manager.clientsMapping[member]
			for _, toClient := range toClients {
				toClient.write(message)
			}
		}
	}

}

type Client struct {
	Id      string
	Socket  *websocket.Conn
	userId  string
	Context echo.Context
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
		var messageRequest model.MessageRequest
		err = json.Unmarshal([]byte(rawMessage), &messageRequest)
		if err != nil {
			log.Println(err.Error())
		}
		var message model.Message = model.Message{
			From:           c.userId,
			ConversationId: messageRequest.ConversationId,
			Message:        messageRequest.Message,
		}
		fmt.Println(message)
		//add message to DB
		err = util.SendMessage(message, c.Context.Get("user").(*jwt.Token))
		if err != nil {
			fmt.Println(err)
		}

		manager.send(message)
	}
}
func (c *Client) write(message model.Message) {
	jsonMessage, _ := json.Marshal(message)
	c.Socket.WriteMessage(websocket.TextMessage, []byte(string(jsonMessage)))
}
func (handler SocketHandler) WsHandler(c echo.Context) error {
	var (
		upgrader = websocket.Upgrader{}
	)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	//creating in memory user-tag to client mapping
	userId := util.GetUserIdFromJWT(c)
	fmt.Println(userId)
	client := &Client{Id: uuid.NewString(), Socket: ws, userId: userId, Context: c}
	manager.clientsMapping[userId] = append(manager.clientsMapping[userId], client)
	manager.DB = handler.DB
	go client.read()
	return nil
}
