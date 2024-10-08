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
	w.WriteHeader(statusCode)
  w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func HandleError(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Println(message, err)
	WriteError(w, err.Error(), statusCode)
	logging.ErrorLogger.Error(message)
}

