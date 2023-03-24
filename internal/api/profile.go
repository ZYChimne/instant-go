package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

type cachedUser struct {
	user      model.User
	timestamp int64
}

func GetUserProfileDetail(c *gin.Context) {
	userID := c.MustGet("UserID")
	targetID := c.Query("userID")
	if targetID == "" {
		targetID = userID.(string)
	}
	errMsg := "Get userinfo error"
	key := strings.Join([]string{"profileDetail", targetID}, ":")
	if val, ok := database.UserCache.Get(key); ok {
		cachedUser := val.(cachedUser)
		if time.Now().Unix()-cachedUser.timestamp < 60 {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": cachedUser.user})
			return
		} else {
			database.UserCache.Remove(key)
		}
	}
	var user model.User
	if userJson, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		log.Println(errMsg, " ", err.Error())
	} else {
		err := json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			log.Println("Unmarshal error ", err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
			return
		}
	}
	err := database.GetUserProfileByID(targetID).Decode(&user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, UndefinedError)
	}
	if redisExpireTime >= 0 {
		if err := database.RedisClient.Set(ctx, key, userJson, redisExpireTime).Err(); err != nil {
			log.Println("Redis error ", err.Error())
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
}

func GetUserProfile(c *gin.Context) {
	userID := c.MustGet("UserID")
	targetID := c.Query("userID")
	if targetID == "" {
		targetID = userID.(string)
	}
	errMsg := "Get userinfo error"
	key := strings.Join([]string{"profile", targetID}, ":")
	var user model.SimpleUser
	if userJson, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
		handleError(c, err, 0, errMsg, Warning)
	} else {
		err := json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			handleError(c, err, 0, "Unmarshal error", Warning)
		} else {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
			return
		}
	}
	err := database.GetUserProfileByID(targetID).Decode(&user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
	}
	if redisExpireTime >= 0 {
		if err := database.RedisClient.Set(ctx, key, userJson, redisExpireTime).Err(); err != nil {
			log.Println("Redis error ", err.Error())
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
}

func QueryUsers(c *gin.Context) {
	userID := c.MustGet("UserID")
	keyword := c.Query("keyword")
	errMsg := "Query users error"
	if keyword == "" {
		handleError(c, nil, http.StatusBadRequest, errMsg, UndefinedError)
		return
	}
	index, err := strconv.ParseInt(c.Query("index"), 0, 64)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, UndefinedError)
		return
	}
	rows, err := database.QueryUsers(userID.(string), keyword, index, pageSize)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	defer rows.Close(ctx)
	users := []model.QueryUser{}
	for rows.Next(ctx) {
		var user model.QueryUser
		err := rows.Decode(&user)
		if err != nil {
			handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
}
