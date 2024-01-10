package model

type Message struct {
	From           string `json:"from" validate:"required"`
	Message        string `json:"message" validate:"required"`
	ConversationId string `json:"conversation_id" bson:"conversation_id" validate:"required"`
}
type MessageRequest struct {
	Message        string `json:"message" validate:"required"`
	ConversationId string `json:"conversation_id" bson:"conversation_id" validate:"required"`
}
