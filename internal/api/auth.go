package api

import (
	"net/http"
	"net/mail"
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "pong"})
}

func Register(c *gin.Context) {
	var userSchema schema.CreateUser
	errMsg := "Register error"
	if err := c.Bind(&userSchema); err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	hash, err := util.HashPassword(userSchema.Password)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	userModel := model.User{
		Email:        userSchema.Email,
		Phone:        userSchema.Phone,
		Password:     hash,
		Username:     userSchema.Username,
		Nickname:     userSchema.Nickname,
		Type:         0,
		Avatar:       userSchema.Avatar,
		Gender:       userSchema.Gender,
		Country:      userSchema.Country,
		State:        userSchema.State,
		City:         userSchema.City,
		ZipCode:      userSchema.ZipCode,
		Birthday:     userSchema.Birthday,
		School:       userSchema.School,
		Company:      userSchema.Company,
		Job:          userSchema.Job,
		MyMode:       userSchema.MyMode,
		Introduction: userSchema.Introduction,
		CoverPhoto:   userSchema.CoverPhoto,
		FollowingCount:   0,
		FollowerCount:    0,
	}
	err = database.CreateUser(&userModel)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Register success"})
}

func GetToken(c *gin.Context) {
	var userSchema schema.LoginUser
	errMsg := "Please check if your account or password is correct"
	if err := c.Bind(&userSchema); err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	if len(userSchema.Email) > 0 {
		if _, err := mail.ParseAddress(userSchema.Email); err != nil {
			handleError(c, err, errMsg, ParameterError)
			return
		}
		var userModel model.User
		if err := database.LoginUserByEmail(userSchema.Email, &userModel); err != nil {
			handleError(c, err, errMsg, DatabaseError)
			return
		}
		if !util.CheckPasswordHash(userSchema.Password, userModel.Password) {
			handleError(c, nil, errMsg, PasswordError)
			return
		}
		token := util.GenerateJwt(userModel.ID)
		c.JSON(http.StatusOK, gin.H{"data": token})
		return
	}
	if len(userSchema.Phone) > 0 {
		for _, char := range userSchema.Phone {
			if char < '0' || char > '9' {
				handleError(c, nil, errMsg, ParameterError)
				return
			}
		}
		var userModel model.User
		if err := database.LoginUserByPhone(userSchema.Phone, &userModel); err != nil {
			handleError(c, err, errMsg, DatabaseError)
			return
		}
		if !util.CheckPasswordHash(userSchema.Password, userModel.Password) {
			handleError(c, nil, errMsg, PasswordError)
			return
		}
		token := util.GenerateJwt(userModel.ID)
		c.JSON(http.StatusOK, gin.H{"data": token})
		return
	}
	handleError(c, nil, errMsg, ParameterError)
}
