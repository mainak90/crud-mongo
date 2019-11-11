package controllers

import (
	"context"
	"crud-mongo/models"
	"crud-mongo/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct{}

var trainer []models.Trainer

func (c Controller) InsertDocument(client *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var trainer models.Trainer
		json.NewDecoder(r.Body).Decode(&trainer)
		param := mux.Vars(r)
		collection := client.Collection(param["collection"])
		insertResult, err := collection.InsertOne(context.TODO(), &trainer)
		utils.LogFatal(err)
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		json.NewEncoder(w).Encode(&trainer)
	}
}
