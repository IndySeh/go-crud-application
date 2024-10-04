package main

import (
	"github.com/IndySeh/go-crud-application/internals/handlers"
	"github.com/IndySeh/go-crud-application/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {

	logError := logging.InitLoggers()
	if logError != nil {
		log.Fatalf("Failed to initialize loggers %v", logError)
	}

	logging.InfoLogger.Info("Server Started")

	err := godotenv.Load()
	if err != nil {
		logging.ErrorLogger.Error("Fail to load .env file")
		log.Fatal("Error in loading .env file. Please check logs: /logs/error.log")
	}

	mux := mux.NewRouter()

	mux.Use(LogIncomingRequestMiddleWare)
	
	mux.HandleFunc("/api/users", handlers.GetAllUsersHandler).Methods("GET")
	mux.HandleFunc("/api/users/{id}", handlers.GetUserHandler).Methods("GET")
	mux.HandleFunc("/api/users", handlers.AddUserHandler).Methods("POST")
	mux.HandleFunc("/api/users", handlers.UpdateUserHandler).Methods("PUT")
	mux.HandleFunc("/api/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	logging.InfoLogger.Info("Running on Port: 8090")
	log.Println("Server is running on Port:8090")
	log.Fatal(http.ListenAndServe(":8090", mux))
}

func LogIncomingRequestMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request Method: %s, Request URL: %s", r.Method, r.URL)
		logging.RequestLogger.Info("Request Method: " + r.Method + " Requested URL: " + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
