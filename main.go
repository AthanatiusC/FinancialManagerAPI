/* Copyright (C) Ahmad Saugi & Lexi Anugrah - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Lexi Anugrah <athanatius@4save.me>, November 2019
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	// "github.com/jinzhu/gorm"

	"github.com/AthanatiusC/FinanceManager/app"
	"github.com/AthanatiusC/FinanceManager/controllers"
)

func main() {
	router := mux.NewRouter()
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode("PREFLIGHT OK")
		})
	APP_PORT := os.Getenv("PORT")
	if APP_PORT == "" {
		APP_PORT = "8088"
	}

	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(app.JwtAuthentication)

	v1Task := apiV1.PathPrefix("/manager").Subrouter()
	v1Task.HandleFunc("/history/{id}", controllers.GetAllTransaction).Methods("GET", "OPTIONS")    // View All
	v1Task.HandleFunc("/detail/{id}", controllers.GetTransactionDetails).Methods("GET", "OPTIONS") // Get Detail
	v1Task.HandleFunc("/insert", controllers.InsertTransaction).Methods("POST", "OPTIONS")         // Insert
	v1Task.HandleFunc("/update", controllers.TransactionUpdate).Methods("PUT", "OPTIONS")          // Update
	v1Task.HandleFunc("/upload", controllers.RecieveImage).Methods("POST", "OPTIONS")              // Upload
	v1Task.HandleFunc("/delete/{id}", controllers.TransactionDelete).Methods("DELETE", "OPTIONS")  // Delete

	v1User := apiV1.PathPrefix("/user").Subrouter()
	v1User.HandleFunc("/register", controllers.UserCreate).Methods("POST", "OPTIONS") // Register
	v1User.HandleFunc("/auth", controllers.Auth).Methods("POST", "OPTIONS")           // Authentication
	v1User.HandleFunc("/{id}", controllers.UserGetOne).Methods("GET", "OPTIONS")      // Get Detail
	v1User.HandleFunc("/update", controllers.UserUpdate).Methods("PUT", "OPTIONS")    // Update
	v1User.HandleFunc("/delete", controllers.UserDelete).Methods("DELETE", "OPTIONS") // Delete

	fmt.Println("App running on port " + APP_PORT)
	log.Fatal(http.ListenAndServe(":"+APP_PORT, router))
}
