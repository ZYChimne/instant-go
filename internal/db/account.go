package database

import (
	"zychimne/instant/pkg/model"
)

const FuzzyMatchThreshold = 2

func CreateAccount(user *model.User) error {
	return PostgresDB.Create(&user).Error
}

func CreateAccounts(users *[]model.User) error {
	return PostgresDB.Create(&users).Error
}

func GetAccountByEmail(email string, user *model.User) error {
	return PostgresDB.Where("email = ?", email).First(&user).Error
}

func GetAccountByPhone(phone string, user *model.User) error {
	return PostgresDB.Where("phone = ?", phone).First(&user).Error
}

func GetAccountByUsername(username string, user *model.User) error {
	return PostgresDB.Where("username = ?", username).First(&user).Error
}

func GetAccountByID(id uint, user *model.User) error {
	return PostgresDB.Where("id = ?", id).First(&user).Error
}

func DeleteAccountByID(id uint) error {
	return PostgresDB.Where("id = ?", id).Delete(&model.User{}).Error
}

func SearchAccounts(keyword string, offset int, limit int, users *[]model.User) error {
	return PostgresDB.Where("difference(nickname, ?) > "+string(rune(FuzzyMatchThreshold)), keyword).Offset(offset).Limit(limit).Find(&users).Error
}

func GetBasicAccount(user *model.BasicUser) error {
	return PostgresDB.Table("users").Select("username", "nickname", "avatar", "user_type").Where("id = ?", user.UserID).Scan(&user).Error
}