package model

type Message struct {
	Id             string `json:"_id" bson:"_id" validate:"required"`
	From           string `json:"from" validate:"required"`
	Message        string `json:"message" validate:"required"`
	ConversationId string `json:"conversation_id" bson:"conversation_id" validate:"required"`
}
