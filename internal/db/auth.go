package database

import (
	"zychimne/instant/pkg/model"
)

func CreateUser(user *model.User) error {
	return PostgresDB.Create(&user).Error
}

func CreateUsers(users *[]model.User) error {
	return PostgresDB.Create(&users).Error
}

func LoginUserByEmail(email string, user *model.User) error {
	return PostgresDB.Select("id", "password").Where("email = ?", email).First(&user).Error
}

func LoginUserByPhone(phone string, user *model.User) error {
	return PostgresDB.Select("id", "password").Where("phone = ?", phone).First(&user).Error
}

func GetUserByID(id string, user *model.User) error {
	return PostgresDB.Where("id = ?", id).First(&user).Error
}
