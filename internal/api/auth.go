package api

import (
	"log"
	"net/http"
	"strings"
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	errMsg := "Register error"
	if err := c.Bind(&user); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	result, err := database.Register(user)
	if err != nil {
		log.Println("Database result:, error: ", result, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusCreated})
}

func GetToken(c *gin.Context) {
	var user model.User
	errMsg := "Please check if your account or password is correct"
	if err := c.Bind(&user); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Get token error"})
		return
	}
	key := strings.Join([]string{"token", user.MailBox, user.Password}, ":")
	if token, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		log.Println("Get token error ", err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": token})
		return
	}
	password := user.Password
	if err := database.GetUser(user.MailBox, &user); err != nil {
		log.Println("Database error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": errMsg})
		return
	}
	if !util.CheckPasswordHash(password, user.Password) {
		log.Println("Password error")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": errMsg})
		return
	}
	token := util.GenerateJwt(user.UserID)
	if err := database.RedisClient.Set(ctx, key, token, 0).Err(); err != nil {
		log.Println("Redis error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": token})
}
