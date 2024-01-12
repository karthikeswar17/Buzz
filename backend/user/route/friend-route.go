package route

import (
	"github.com/karthikeswar17/user/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func FriendRoute(e *echo.Echo, DB *mongo.Client) {
	var friendHandler handler.FriendHandler = handler.FriendHandler{DB: *DB}
	e.GET("friend/list", friendHandler.GetFriendList)
	e.POST("friend/request", friendHandler.SendFriendRequest)
	e.GET("friend/request", friendHandler.GetFriendRequests)
	e.GET("friend/request/:id/accept", friendHandler.AcceptFriendRequest)
	e.GET("friend/request/:id/decline", friendHandler.DeclineFriendRequest)
	e.GET("friend/:id/conversation", friendHandler.DeclineFriendRequest)
}
