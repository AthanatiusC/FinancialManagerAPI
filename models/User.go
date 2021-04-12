package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Token struct {
	UserID uint
	jwt.StandardClaims
}

//User structure
type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id" `
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password" `
	FullName     string             `json:"fullname" bson:"fullname"`
	DoB          primitive.DateTime `json:"dob" bson:"dob"`
	Address      string             `json:"address" bson:"address"`
	AccessToken  string             `json:"access_token" bson:"access_token" `
	RefreshToken string             `json:"refresh_token" bson:"refresh_token" `
}
