package main

import (
	"demo/package/dto"
	"demo/package/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// This is a Controller, all the incoming requests are delegated to Service layer
// and errors from Service layer are handled

// Returns all the Users present in the database
func (demo *DemoApp) getAllUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	// Get All Users from Service
	//var users []dto.UserDTO = service.GetAllUsers()
	var users []dto.UserDTO = service.GetAllUsersFromDB()

	if users == nil {
		http.Error(writer, "Empty Database!", http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

// Returns a User based on incoming user_id
func (demo *DemoApp) getUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userId := vars["user_id"]

	// If not a valid user_id, error out the request
	if userId == "" {
		http.Error(writer, "Invalid User Id", http.StatusBadRequest)
		return
	}

	// Get User from Service by passing the user_id
	//user := service.GetUser(userId)
	user := service.GetUserFromDB(userId)
	if user == nil {
		http.Error(writer, "Unable to find User for user_id"+userId, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

// Create a User based on request body
func (demo *DemoApp) createUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user dto.UserDTO
	_ = json.NewDecoder(request.Body).Decode(&user)

	// Create User by calling the Service
	//createdUser := service.CreateUser(user)
	createdUser := service.CreateUserFromDB(user)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(createdUser)
}

// Update the User based on the request body matching user_id
func (demo *DemoApp) updateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userId := vars["user_id"]
	var user dto.UserDTO
	_ = json.NewDecoder(request.Body).Decode(&user)

	// If not a valid user_id, error out the request
	if userId == "" {
		http.Error(writer, "Invalid User Id", http.StatusBadRequest)
		return
	}

	// Update User and get Updated User from Service by passing the user_id
	//updatedUser := service.UpdateUser(userId, user)
	updatedUser := service.UpdateUserFromDB(userId, user)

	if updatedUser == nil {
		http.Error(writer, "Unable to find User for user_id"+userId, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(updatedUser)
}

// Delete the User based on user_id
func (demo *DemoApp) deleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	userId := vars["user_id"]

	// If not a valid user_id, error out the request
	if userId == "" {
		http.Error(writer, "Invalid User Id", http.StatusBadRequest)
		return
	}

	// Delete User from Users in Service by passing the user_id
	//deletedUser := service.DeleteUser(userId)
	deletedUser := service.DeleteUserFromDB(userId)
	if deletedUser == nil {
		http.Error(writer, "Unable to find User for user_id"+userId, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(deletedUser)

}
