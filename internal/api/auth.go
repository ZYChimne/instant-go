package api

import (
	"net/http"
	"strings"
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "pong"})
}

func Register(c *gin.Context) {
	var user model.User
	errMsg := "Register error"
	if err := c.Bind(&user); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	_, err := database.Register(user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusCreated})
}

func GetToken(c *gin.Context) {
	var user model.User
	errMsg := "Please check if your account or password is correct"
	if err := c.Bind(&user); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	key := strings.Join([]string{"token", user.MailBox, user.Password}, ":")
	if token, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		handleError(nil, err, 0, errMsg, RedisError)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": token})
		return
	}
	password := user.Password
	if err := database.GetUser(user.MailBox, &user); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	if !util.CheckPasswordHash(password, user.Password) {
		handleError(c, nil, http.StatusBadRequest, errMsg, PasswordError)
		return
	}
	token := util.GenerateJwt(user.UserID)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": token})
	putInRedis(key, token)
}
