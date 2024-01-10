package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/karthikeswar17/user/model"
	"github.com/karthikeswar17/user/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FriendHandler struct {
	DB mongo.Client
}

func (handler *FriendHandler) GetFriendList(c echo.Context) error {
	var err error
	var userTag string
	var user model.User
	var searchQuery string = c.QueryParam("q")
	var friendsResponse []model.FriendListResponse = []model.FriendListResponse{}
	var userCollection = util.GetCollection(&handler.DB, "User")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, userTag = util.GetEmailAndTagFromJWT(c)
	query := bson.M{"tag": userTag}
	if searchQuery != "" {
		query["name"] = bson.M{"$regex": searchQuery, "$options": "i"}
	}

	err = userCollection.FindOne(ctx, query).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "user not found"})
	}
	if user.Friends != nil {
		opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "tag", Value: 1}, {Key: "name", Value: 1}, {Key: "email", Value: 1}})
		cursor, err := userCollection.Find(ctx, bson.M{"_id": bson.M{"$in": user.Friends}}, opts)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
		}
		err = cursor.All(context.TODO(), &friendsResponse)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
		}
	}

	return c.JSON(http.StatusOK, friendsResponse)

}

func (handler *FriendHandler) SendFriendRequest(c echo.Context) error {
	var err error
	var friendRequestRequest model.FriendRequestRequest
	var friendCollection = util.GetCollection(&handler.DB, "Friend")
	err = (&echo.DefaultBinder{}).BindBody(c, &friendRequestRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed"})
	}
	//get current user
	var sender model.User
	var userTag string
	var userCollection = util.GetCollection(&handler.DB, "User")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, userTag = util.GetEmailAndTagFromJWT(c)
	err = userCollection.FindOne(ctx, bson.M{"tag": userTag}).Decode(&sender)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}

	//get reciver id
	var receiver model.User
	err = userCollection.FindOne(ctx, bson.M{"tag": friendRequestRequest.Tag}).Decode(&receiver)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "Given Friend Tag not found"})
	}
	// Duplicate Validation
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = friendCollection.FindOne(ctx, bson.M{"sender": sender.Id, "receiver": receiver.Id}).Err()
	if err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "Duplicate Request"})
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//vaidation
	if sender.Id == receiver.Id {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Success", "message": "self friend request cannot be sent"})
	}
	var friendRequest = model.FriendRequest{
		Id:       uuid.NewString(),
		Sender:   sender.Id,
		Receiver: receiver.Id,
	}

	_, err = friendCollection.InsertOne(ctx, friendRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "user not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "Success"})

}

func (handler *FriendHandler) GetFriendRequests(c echo.Context) error {
	var err error
	var userTag string
	var user model.User
	var friendRequestList []model.FriendRequest
	var userList []model.User = []model.User{}
	var friendRequestListResponse []model.FriendRequestListResponseItem = []model.FriendRequestListResponseItem{}

	//get user
	var userCollection = util.GetCollection(&handler.DB, "User")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, userTag = util.GetEmailAndTagFromJWT(c)
	err = userCollection.FindOne(ctx, bson.M{"tag": userTag}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "user not found"})
	}
	var friendCollection = util.GetCollection(&handler.DB, "Friend")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//received
	cursor, err := friendCollection.Find(ctx, bson.M{"receiver": user.Id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	err = cursor.All(context.TODO(), &friendRequestList)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	var senderIds []string
	var senderIdsMap = map[string]string{}
	for _, friendRequest := range friendRequestList {
		senderIds = append(senderIds, friendRequest.Sender)
		senderIdsMap[friendRequest.Sender] = friendRequest.Id
	}
	//get sender details
	if senderIds != nil {
		opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "tag", Value: 1}, {Key: "name", Value: 1}, {Key: "email", Value: 1}})
		cursor, err = userCollection.Find(ctx, bson.M{"_id": bson.M{"$in": senderIds}}, opts)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
		}
		err = cursor.All(context.TODO(), &userList)
	}

	for _, user := range userList {
		friendRequestListResponse = append(friendRequestListResponse, model.FriendRequestListResponseItem{
			Id:            user.Id,
			Email:         user.Email,
			Name:          user.Name,
			Tag:           user.Tag,
			FriendRequest: senderIdsMap[user.Id],
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	return c.JSON(http.StatusOK, friendRequestListResponse)

}
func (handler *FriendHandler) AcceptFriendRequest(c echo.Context) error {
	var err error
	var userTag string
	var user model.User
	var friendRequestId string = c.Param("id")
	fmt.Println(friendRequestId)
	var friendRequest model.FriendRequest

	//get user
	var userCollection = util.GetCollection(&handler.DB, "User")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, userTag = util.GetEmailAndTagFromJWT(c)
	err = userCollection.FindOne(ctx, bson.M{"tag": userTag}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "user not found"})
	}
	var friendCollection = util.GetCollection(&handler.DB, "Friend")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = friendCollection.FindOne(ctx, bson.M{"_id": friendRequestId}).Decode(&friendRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "Unable to find the given friend request id"})
	}
	if friendRequest.Receiver != user.Id {
		return c.JSON(http.StatusForbidden, echo.Map{"status": "Failed", "message": "Not authorized to accpet others friend request"})
	}

	//sender
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": friendRequest.Sender}, bson.M{"$push": bson.M{"friends": friendRequest.Receiver}})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}

	//reciver
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": friendRequest.Receiver}, bson.M{"$push": bson.M{"friends": friendRequest.Sender}})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}

	//delete request
	_, err = friendCollection.DeleteOne(ctx, bson.M{"_id": friendRequestId})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "Success"})
}
func (handler *FriendHandler) DeclineFriendRequest(c echo.Context) error {
	var err error
	var userTag string
	var user model.User
	var friendRequestId string = c.Param("id")

	//get user
	var userCollection = util.GetCollection(&handler.DB, "User")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, userTag = util.GetEmailAndTagFromJWT(c)
	err = userCollection.FindOne(ctx, bson.M{"tag": userTag}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "user not found"})
	}
	var friendCollection = util.GetCollection(&handler.DB, "Friend")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := friendCollection.DeleteOne(ctx, bson.M{"_id": friendRequestId, "receiver": user.Id})
	fmt.Println(bson.M{"_id": friendRequestId, "receiver": user.Id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": err})
	}
	if result.DeletedCount == 0 {
		return c.JSON(http.StatusForbidden, echo.Map{"status": "Failed", "message": "Check if the given request exists/ you enough have previlage to decline request"})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "Success"})
}
