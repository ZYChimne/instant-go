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
	err := database.GetUserInfo(&user)
	if err != nil {
		log.Panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
}
