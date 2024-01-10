package model

type FriendRequest struct {
	Id       string `json:"_id" bson:"_id" validate:"required"`
	Sender   string `json:"sender" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
}

type FriendRequestRequest struct {
	Tag string `json:"tag" bson:"_id" validate:"required"`
}
type FriendRequestListResponseItem struct {
	Id            string `json:"_id" bson:"_id" validate:"required"`
	Tag           string `json:"tag" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	FriendRequest string `json:"friend_request" validate:"required"`
}

type FriendListResponse struct {
	Id    string `json:"_id" bson:"_id" validate:"required"`
	Tag   string `json:"tag" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}
