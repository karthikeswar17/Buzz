package model

type User struct {
	Id           string   `json:"_id" bson:"_id" validate:"required"`
	Tag          string   `json:"tag" validate:"required"`
	Name         string   `json:"name" validate:"required"`
	Email        string   `json:"email" validate:"required"`
	PasswordHash string   `json:"password_hash" validate:"required"`
	Friends      []string `json:"friends" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Id       string `json:"_id" bson:"_id" validate:"required"`
	Tag      string `json:"tag" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
