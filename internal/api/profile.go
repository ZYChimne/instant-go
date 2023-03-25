package api

import (
	"net/http"
	"strconv"
	"strings"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetUserProfileDetail(c *gin.Context) {
	userID := c.MustGet("UserID")
	targetID := c.Query("userID")
	if targetID == "" {
		targetID = userID.(string)
	}
	errMsg := "Get userinfo error"
	key := strings.Join([]string{"profileDetail", targetID}, ":")
	if getFromCache(c, key, database.UserCache) {
		return
	}
	var user model.User
	err := database.GetUserProfileByID(targetID).Decode(&user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
	putInCache(key, user, database.UserCache)
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
	if getFromCache(c, key, database.SimpleUserCache) {
		return
	}
	err := database.GetUserProfileByID(targetID).Decode(&user)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": user})
	putInCache(key, user, database.SimpleUserCache)
}

func QueryUsers(c *gin.Context) {
	userID := c.MustGet("UserID")
	keyword := c.Query("keyword")
	errMsg := "Query users error"
	if keyword == "" {
		handleError(c, nil, http.StatusBadRequest, errMsg, BindError)
		return
	}
	index, err := strconv.ParseInt(c.Query("index"), 0, 64)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
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
