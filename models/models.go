package models

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// username usera
// password b8VEuQAyiVocR8RQ

// username client
// password H76dQCR4DZ8bhBLI
func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://client:H76dQCR4DZ8bhBLI@finagecluster0.lv8aj.mongodb.net/FinanceManager?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	db = client.Database("FinanceManager")
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")

}

func GetDB(whatDB string) *mongo.Database {
	if whatDB == "main" {
		return db
	} else {
		return db
	}
	return db
}
