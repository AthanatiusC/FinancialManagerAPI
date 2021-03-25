package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Task structure
type Transaction struct {
	ID          primitive.ObjectID `json:"id" bson:"_id" `
	UserID      primitive.ObjectID `json:"userid" bson:"uid"`
	Transaction string             `json:"transaction" bson:"transaction"`
	Recipt      string             `json:"recipt" bson:"recipt"`
	Description string             `json:"description" bson:"description"`
	Type        string             `json:"type" bson:"type"`
	Time        time.Time          `json:"time" bson:"time"`
}
