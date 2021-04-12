package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AthanatiusC/FinanceManager/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TaskGetAll for return index api
func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	raw := mux.Vars(r)["id"]
	userid, _ := primitive.ObjectIDFromHex(raw)
	transactions := []models.Transaction{}
	transaction := models.Transaction{}

	ctx := context.TODO() // Options to the database.
	coll, err := models.GetDB("main").Collection("transactions").Find(ctx, bson.M{"userid": userid})
	if err != nil {
		fmt.Println(err)
	}

	for coll.Next(ctx) {
		coll.Decode(&transaction)
		// transaction.Time = transaction.Time
		transactions = append(transactions, transaction)
		transaction = models.Transaction{}
	}
	respondJSON(w, 200, "Success get all transaction history for current users!", transactions)
}

// TaskGetOne for returning single item
func GetTransactionDetails(w http.ResponseWriter, r *http.Request) {
	raw := mux.Vars(r)["id"]
	taskID, _ := primitive.ObjectIDFromHex(raw)
	var manager models.Transaction
	err := models.GetDB("main").Collection("transactions").FindOne(context.TODO(), bson.M{"_id": taskID}).Decode(&manager)

	if err != nil {
		fmt.Println(err)
		respondJSON(w, 200, "Transaction not found!", map[string]interface{}{})
		return
	}
	respondJSON(w, 200, "Get Transaction Detail", manager)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	var manager models.Transaction
	json.NewDecoder(r.Body).Decode(&manager)
	manager.ID = primitive.NewObjectID()
	models.GetDB("main").Collection("transactions").InsertOne(context.TODO(), &manager)
	respondJSON(w, 200, "Success Create New Task!", manager)
}

func TransactionDelete(w http.ResponseWriter, r *http.Request) {
	raw := mux.Vars(r)["id"]
	taskid, _ := primitive.ObjectIDFromHex(raw)

	deleteResult, err := models.GetDB("main").Collection("transactions").DeleteOne(context.TODO(), bson.M{"_id": taskid})
	if err != nil {
		respondJSON(w, 404, "Error!", err)
		return
	}
	respondJSON(w, 200, "manager deleted", deleteResult)
}

func TransactionUpdate(w http.ResponseWriter, r *http.Request) {
	var manager models.Transaction
	json.NewDecoder(r.Body).Decode(&manager)
	res, err := models.GetDB("main").Collection("transactions").UpdateOne(context.TODO(), bson.M{"_id": manager.ID, "userid": manager.UserID}, bson.D{{Key: "$set", Value: manager}})
	if err != nil {
		respondJSON(w, 404, "Error occured", err)
		return
	}
	respondJSON(w, 200, "manager Updated!", res)
}
