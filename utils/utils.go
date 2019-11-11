package utils

import (
	"encoding/json"
	"go-crud-jwt/models"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
