package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userID := c.MustGet("UserID")
	targetID := c.Query("userID")
	if targetID == "" {
		targetID = userID.(string)
	}
	errMsg := "Get userinfo error"
	key := strings.Join([]string{"profile", targetID}, ":")
	var user model.User
	if cache, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		log.Println(errMsg, " ", err.Error())
	} else {
		err := json.Unmarshal([]byte(cache), &user)
		if err != nil {
			log.Println("Unmarshal error ", err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
			return
		}
	}
	user.UserID = targetID
	err := database.GetUserInfo(&user)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	cache, err := json.Marshal(user)
	if err != nil {
		log.Println("Marshal error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
	}
	if err := database.RedisClient.Set(ctx, key, cache, redisExpireTime).Err(); err != nil {
		log.Println("Redis error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
}
