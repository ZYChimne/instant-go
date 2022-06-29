package database

import (
	"database/sql"
	"log"
	"time"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"
)

func Register(user model.User) (sql.Result, error){
	hash, err :=util.HashPassword(user.Password)
	if err != nil {
		log.Fatal("password hash error", err.Error())
	}
	query := `INSERT INTO accounts (mailbox, phone, pass_word, username, create_time, update_time, avatar, gender, country, province, city, birthday, school, company, job, introduction, profile_image, tag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	return db.Exec(query, user.MailBox, user.Phone, hash, user.Username, time.Now(), time.Now(), 0, user.Gender, user.Country, user.Province, user.City, user.Birthday, user.School, user.Company, user.Job, user.Introduction, 0, user.Tag)
}

func GetToken(user model.User, userID *int, hash *string) error{
	query := `SELECT user_id, pass_word FROM accounts WHERE mailbox = ?`
	return db.QueryRow(query, user.MailBox).Scan(&userID, &hash);
}