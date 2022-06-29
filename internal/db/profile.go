package database

import (
	"zychimne/instant/pkg/model"
)

func GetUserInfo(user *model.User) (error){
	query := "SELECT username, create_time, avatar, gender, country, province, city, birthday, school, company, my_mode, job, introduction, cover, tag FROM accounts WHERE user_id = ?"
	return db.QueryRow(query, 14).Scan(&user.Username, &user.CreateTime, &user.Avatar, &user.Gender, &user.Country, &user.Province, &user.City, &user.Birthday, &user.School, &user.Company, &user.MyMode, &user.Job, &user.Introduction, &user.CoverPhoto, &user.Tag)
}