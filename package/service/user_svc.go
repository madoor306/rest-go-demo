package service

import (
	dto "demo/package/dto"
	"math/rand"
	"strconv"
)

//Service layer which interacts with the Repository

// Using slice instead of Repo layer as of now - TODO
var users []dto.UserDTO

// Get All Users from repository
func GetAllUsers() []dto.UserDTO {

	if users == nil || len(users) == 0 {
		return nil
	}

	return users
}

// Get User by user_id
func GetUser(user_id string) *dto.UserDTO {

	//Loop through all the Users, if found return
	for _, item := range users {

		if item.UserId == user_id {
			return &item
		}
	}

	return nil
}

// Create User in the Repository
func CreateUser(user dto.UserDTO) *dto.UserDTO {
	user.UserId = strconv.Itoa(rand.Intn(100000000))
	user.Address.AddressId = strconv.Itoa(rand.Intn(100000000))
	users = append(users, user)
	return &user
}

// Update User in the Repository
func UpdateUser(userId string, user dto.UserDTO) *dto.UserDTO {

	for index, item := range users {

		if item.UserId == userId {
			users = append(users[:index], users[index+1:]...)
			user.UserId = userId
			users = append(users, user)
			return &user
		}
	}

	return nil
}

// Delete User from Repository
func DeleteUser(user_id string) *dto.UserDTO {

	for index, item := range users {

		if item.UserId == user_id {
			users = append(users[:index], users[index+1:]...)
			return &item
		}
	}

	return nil
}
