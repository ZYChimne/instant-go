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
		log.Fatal(err.Error())
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": user})
}
