package api

import (
	"log"
	"net/http"
	"zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var user model.User
	query := "SELECT username, create_time, avatar, gender, country, province, city, birthday, school, company, my_mode, job, introduction, cover, tag FROM accounts WHERE user_id = ?"
	db := database.ConnectDatabase()
	err := db.QueryRow(query, 14).Scan(&user.Username, &user.CreateTime, &user.Avatar, &user.Gender, &user.Country, &user.Province, &user.City, &user.Birthday, &user.School, &user.Company, &user.MyMode, &user.Job, &user.Introduction, &user.CoverPhoto, &user.Tag)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
