package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	controllers "go-test-program/controllers"
	u "go-test-program/utils"
	"log"
	"net/http"
	"os"
)

func HandleRequests() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_, err := u.GetSqlConnection()
	if err != nil {
		log.Println(err.Error())
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/messages", controllers.GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", controllers.AddMessage).Methods("POST")
	router.HandleFunc("/api/messages/{id}", controllers.GetMessage).Methods("GET")
	router.HandleFunc("/api/messages/{id}", controllers.DeleteMessage).Methods("DELETE")
	router.HandleFunc("/api/messages/{id}", controllers.UpdateMessage).Methods("PUT")

	fmt.Println("Listening port:", port)
	err = http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	HandleRequests()
}
