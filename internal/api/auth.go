package api

import (
	"log"
	"net/http"
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": "400", "errMsg": err.Error()})
	}
	result, err := database.Register(user)
	if err != nil {
		log.Panic("database result:, error: ", result, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusCreated})
}

func GetToken(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
	}
	var (
		userID int
		hash   string
	)
	if err := database.GetToken(user, &userID, &hash); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": http.StatusForbidden, "message": "Please check if your account or password is correct"})
		log.Println("database error: ", err.Error(), "& account not found")
		return
	}
	if !util.CheckPasswordHash(user.Password, hash) {
		log.Println("password error")
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": gin.H{"token": util.GenerateJwt(userID)}, "message": "ok"})
}
