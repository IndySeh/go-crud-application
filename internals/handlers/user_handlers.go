package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/IndySeh/go-crud-application/internals/db"
	"github.com/IndySeh/go-crud-application/internals/repository"
	"github.com/IndySeh/go-crud-application/internals/utils"
	"github.com/IndySeh/go-crud-application/pkg/types"
	"github.com/gorilla/mux"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		utils.HandleError(w, err, "error connecting database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	users, err := repository.FetchUsersFromDB(db)

	if err != nil {
		utils.HandleError(w, err, "error in fetching user from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		utils.HandleError(w, err, "error connecting database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	vars := mux.Vars(r) // Get the Map of variable from the request
	idStr := vars["id"]

	userId, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(w, err, "string conversion error", http.StatusInternalServerError)
		return
	}

	user, err := repository.FetchUserFromDB(db, userId)
	if err != nil {
		utils.HandleError(w, err, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		utils.HandleError(w, err, "error connecting database", http.StatusInternalServerError)
	}

	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HandleError(w, err, "string conversion error", http.StatusInternalServerError)
		return
	}

	err = repository.DeleteUserFromDB(db, userID)

	if err != nil {
		utils.HandleError(w, err, "error in deleting user from database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("user deleted successfully")
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		utils.HandleError(w, err, "error connecting database", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	user := &types.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError(w, err, "invalid request body", http.StatusBadRequest)
		return
	}

	err = repository.InsertUserInDB(db, user.Name, user.Email)
	if err != nil {
		utils.HandleError(w, err, "error in inserting user", http.StatusInternalServerError)
		return
	}

	response := types.PostResponse{
		Message: "User created successfully",
		User:    user,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
