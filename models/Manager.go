package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Task structure
type Transaction struct {
	ID          primitive.ObjectID `json:"id" bson:"_id" `
	UserID      primitive.ObjectID `json:"userid" bson:"userid"`
	Transaction string             `json:"transaction" bson:"transaction"`
	Amount      string             `json:"amount" bson:"amount"`
	Recipt      string             `json:"recipt" bson:"recipt"`
	Description string             `json:"description" bson:"description"`
	Types       string             `json:"types" bson:"types"`
	Time        time.Time          `json:"time" bson:"time"`
}
