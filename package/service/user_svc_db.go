package service

import (
	dto "demo/package/dto"
	repo "demo/package/repository"
	"fmt"
)

//Service layer which interacts with the Repository

// Get All Users from repository
func GetAllUsersFromDB() []dto.UserDTO {
	usersDB := repo.DB.GetAllUsersFromRepo()

	if usersDB == nil || len(usersDB) == 0 {
		return nil
	}

	return usersDB
}

// Get User by user_id
func GetUserFromDB(user_id string) *dto.UserDTO {

	return repo.GetUserFromRepo(user_id)

}

// Create User in the Repository
func CreateUserFromDB(user dto.UserDTO) *dto.UserDTO {
	error := repo.DB.CreateUser(&user)
	if error != nil {
		fmt.Println("Invalid User!")
	}
	return &user
}

// Update User in the Repository
func UpdateUserFromDB(user_Id string, user dto.UserDTO) *dto.UserDTO {

	error := repo.DB.UpdateUser(user_Id, &user)
	if error != nil {
		fmt.Println("Unable to Update User!")
	}
	return &user
}

// Delete User from Repository
func DeleteUserFromDB(user_id string) *dto.UserDTO {

	error := repo.DB.DeleteUser(user_id)
	if error != nil {
		fmt.Println("Unable to Delete User!")
		return nil
	}

	return nil
}
