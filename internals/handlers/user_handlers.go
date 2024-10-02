package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/IndySeh/go-crud-application/internals/db"
	"github.com/IndySeh/go-crud-application/internals/repository"
	"github.com/IndySeh/go-crud-application/internals/utils"
	"github.com/IndySeh/go-crud-application/pkg/logging"
	"github.com/IndySeh/go-crud-application/pkg/types"
	"github.com/gorilla/mux"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		log.Println("Error is connecting database", err)
		utils.WriteError(w, "database connection error", http.StatusInternalServerError)
		logging.ErrorLogger.Error("Error in connecting database")
		return
	}

	defer db.Close()

	users, err := repository.FetchUsersFromDB(db)

	if err != nil {
		log.Fatal("Error fetching users")
		utils.WriteError(w, "error in  fetching user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		log.Println("Error connecting Database", err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
	}

	defer db.Close()

	vars := mux.Vars(r) // Get the Map of variable from the request
	idStr := vars["id"]

	userId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "string conversion error", http.StatusInternalServerError)
		return
	}

	user, err := repository.FetchUserFromDB(db, userId)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		log.Println("error connecting database", err)
		http.Error(w, "database error", http.StatusInternalServerError)
	}

	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "string conversion error"}`, http.StatusInternalServerError)
		return
	}

	err = repository.DeleteUserFromDB(db, userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("user deleted successfully")
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		log.Println("error connecting database", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	user := &types.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = repository.InsertUserInDB(db, user.Name, user.Email)
	if err != nil {
		log.Println("Error in inserting record: ", err)
		http.Error(w, `{"error":"error in inserting user"}`, http.StatusInternalServerError)
		return
	}

	response := types.PostResponse{
		Message: "User created successfully",
		User:    user,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
