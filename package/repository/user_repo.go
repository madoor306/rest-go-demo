package repository

import (
	"fmt"

	"demo/app/main"
	dto "demo/package/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbConn() *gorm.DB {
	dsn := "nfcpay:nfcpay.123@tcp(127.0.0.1:3306)/rest-go-demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to Connect the Database - ", err)
	}

	fmt.Println("Connection to Database Successful!", db)
	return db
}

func (repo *main.DemoApp) CreateUser(user *dto.UserDTO) error {
	return repo.DB.Create(user).Error
}

func (repo *main.DemoApp) UpdateUser(user_id string, user *dto.UserDTO) error {
	var userVar dto.UserDTO
	repo.DB.Where("UserId", user_id).Find(&userVar)
	userVar.Age = user.Age
	userVar.FirstName = user.FirstName
	userVar.LastName = user.LastName
	userVar.Address = user.Address

	return repo.DB.Save(&user).Error
}

func (repo *main.DemoApp) DeleteUser(user_id string) error {
	var user dto.UserDTO
	repo.DB.Where("UserId", user_id).Find(&user)
	return repo.DB.Delete(&user).Error
}

func (repo *main.DemoApp) GetUserFromRepo(user_id string) dto.UserDTO {
	var user dto.UserDTO
	repo.DB.Where("UserId", user_id).Find(&user)
	return user
}

func (repo *main.DemoApp) GetAllUsersFromRepo() []dto.UserDTO {
	var users []dto.UserDTO
	repo.DB.Find(&users)

	return users
}
