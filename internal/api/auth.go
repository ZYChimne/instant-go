package api

import (
	"log"
	"net/http"
	"time"
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": "400", "errMsg": err.Error()})
	}
	hash, err := hashPassword(user.Password)
	if err != nil {
		log.Fatal("password hash error", err.Error())
	}
	query := `INSERT INTO accounts (mailbox, phone, pass_word, username, create_time, update_time, avatar, gender, country, province, city, birthday, school, company, job, introduction, profile_image, tag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	db := database.ConnectDatabase()
	result, err := db.Exec(query, user.MailBox, user.Phone, hash, user.Username, time.Now(), time.Now(), 0, user.Gender, user.Country, user.Province, user.City, user.Birthday, user.School, user.Company, user.Job, user.Introduction, 0, user.Tag)
	db.Close()
	if err != nil {
		log.Fatal("database result:, error: ", result, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetToken(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
	log.Println(user);
	var (
		userID int
		hash   string
	)
	query := `SELECT user_id, pass_word FROM accounts WHERE mailbox = ?`
	db := database.ConnectDatabase()
	if err := db.QueryRow(query, user.MailBox).Scan(&userID, &hash); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": "Please check if your account or password is correct"})
		log.Println("database error: ", err.Error(), "& account not found")
		db.Close();
		return
	}
	db.Close()
	if !checkPasswordHash(user.Password, hash) {
		log.Println("password error")
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": gin.H{"token": utilAuth.GenerateJwt(userID)}, "message": "ok"})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
