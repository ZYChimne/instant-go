package database

import (
	"strconv"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"
)

const FuzzyMatchThreshold int = 2

func CreateAccount(user *model.User) error {
	return PostgresDB.Create(&user).Error
}

func CreateAccounts(users *[]model.User) error {
	return PostgresDB.Create(&users).Error
}

func GetAccountByEmail(email string, user *schema.AccountResponse) error {
	return PostgresDB.Table("users").Where("email = ?", email).Order("id").Limit(1).Scan(&user).Error
}

func GetAccountByPhone(phone string, user *schema.AccountResponse) error {
	return PostgresDB.Table("users").Where("phone = ?", phone).Order("id").Limit(1).Scan(&user).Error
}

func GetAccountByUsername(username string, user *schema.AccountResponse) error {
	return PostgresDB.Table("users").Where("username = ?", username).Order("id").Limit(1).Scan(&user).Error
}

func GetAccountByID(id uint, user *schema.AccountResponse) error {
	return PostgresDB.Table("users").Where("id = ?", id).Scan(&user).Error
}

func DeleteAccountByID(user *model.User) error {
	return PostgresDB.Delete(&user).Error
}

func SearchAccounts(keyword string, offset int, limit int, users *[]schema.BasicAccountResponse) error {
	return PostgresDB.Table("users").Where("difference(nickname, ?) > "+strconv.Itoa(FuzzyMatchThreshold), keyword).Offset(offset).Limit(limit).Scan(&users).Error
}

func GetBasicAccount(user *model.BasicUser) error {
	return PostgresDB.Table("users").Select("username", "nickname", "avatar", "user_type").Where("id = ?", user.UserID).Scan(&user).Error
}

func CheckIfAccountExistsByEmail(email string, exists *int64) error {
	return PostgresDB.Table("users").Where("email = ?", email).Order("id").Limit(1).Count(exists).Error
}

func CheckIfAccountExistsByPhone(phone string, exists *int64) error {
	return PostgresDB.Table("users").Where("phone = ?", phone).Order("id").Limit(1).Count(exists).Error
}

func CheckIfAccountExistsByUsername(username string, exists *int64) error {
	return PostgresDB.Table("users").Where("username = ?", username).Order("id").Limit(1).Count(exists).Error
}
