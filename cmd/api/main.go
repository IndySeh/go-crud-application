package main

import (
	"log"
	"net/http"

	"github.com/IndySeh/go-crud-application/internals/db"
	"github.com/IndySeh/go-crud-application/internals/handlers"
	"github.com/IndySeh/go-crud-application/pkg/logging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	router := mux.NewRouter()

	router.HandleFunc("/api/users", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/users", handlers.AddUserHandler).Methods("POST")
	router.HandleFunc("/api/users", UpdateUserHandler).Methods("PATCH")
	router.HandleFunc("/api/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	logging.InfoLogger.Info("Running on Port: 8090")
	log.Println("Server is running on Port:8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		log.Println("Error in connection database", err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		logging.ErrorLogger.Error("Error in Connection database")
		return
	}

	defer db.Close()

	// UserId, err := strconv.Atoi(idStr)

}
