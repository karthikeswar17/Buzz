package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/karthikeswar17/user/model"
	"github.com/karthikeswar17/user/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DB mongo.Client
}

func (handler *UserHandler) Login(c echo.Context) error {
	var err error
	// payload extraction
	var loginRequest model.LoginRequest
	err = (&echo.DefaultBinder{}).BindBody(c, &loginRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	defer cancel()
	var userCollection = util.GetCollection(&handler.DB, "User")
	err = userCollection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "user not found"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed", "message": "Password Mismatch"})
	}
	token, err := util.CreateJWT(user.Id, user.Email, user.Tag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "Failed", "message": "Unable to generate JWT. Please try again"})
	}
	util.SetCookies(c, map[string]string{"token": token, "tag": user.Tag, "email": user.Email, "id": user.Id})
	return c.JSON(http.StatusOK, echo.Map{"status": "Success", "token": token})
}
func (handler *UserHandler) Register(c echo.Context) error {
	var err error
	var registerRequest model.RegisterRequest
	err = (&echo.DefaultBinder{}).BindBody(c, &registerRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "Failed"})
	}
	//hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(string(hashedPassword))
	//creating user struct
	var user model.User = model.User{
		Id:           uuid.NewString(),
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		Tag:          registerRequest.Tag,
		PasswordHash: string(hashedPassword),
	}

	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result *mongo.InsertOneResult
	var userCollection = util.GetCollection(&handler.DB, "User")
	result, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(result)

	return c.JSON(http.StatusOK, echo.Map{"status": "Success"})
}
