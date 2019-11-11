package main

import (
	"log"
	"net/http"

	"crud-mongo/controllers"
	"crud-mongo/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database

func init() {
	gotenv.Load("go.env")
}

func main() {
	client = driver.ConnectMongo()
	router := mux.NewRouter()
	controller := controllers.Controller{}
	router.HandleFunc("/insert/{collection}", controller.InsertDocument(client)).Methods("POST")
	log.Println("Starting http server listeing on port 8902...")
	log.Fatal(http.ListenAndServe(":8902", router))
}
