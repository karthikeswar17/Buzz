package model

type Conversation struct {
	Id      string   `json:"_id" bson:"_id" validate:"required"`
	Members []string `json:"members" validate:"required"`
}
