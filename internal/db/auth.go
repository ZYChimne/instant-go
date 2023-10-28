package database

import (
	"zychimne/instant/pkg/model"
)

func LoginUserByEmail(email string, user *model.User) error {
	return PostgresDB.Select("id", "password").Where("email = ?", email).First(&user).Error
}

func LoginUserByPhone(phone string, user *model.User) error {
	return PostgresDB.Select("id", "password").Where("phone = ?", phone).First(&user).Error
}

func LoginUserByUsername(username string, user *model.User) error {
	return PostgresDB.Select("id", "password").Where("username = ?", username).First(&user).Error
}