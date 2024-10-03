package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IndySeh/go-crud-application/pkg/logging"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func HandleError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Println(message, err)
	WriteError(w, message, statusCode)
	logging.ErrorLogger.Error(message)
}
