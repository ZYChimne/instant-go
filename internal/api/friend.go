package api

import (
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetFriends(c *gin.Context) {
	userID := c.MustGet("userID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if userID != "" && err == nil {
		users := make([]model.User, pageSize)
		rows, err := database.GetPotentialFollowing(userID.(string), index, pageSize)
		if err != nil {
			log.Panic(err)
		}
		defer rows.Close(ctx)
		cnt := 0
		for rows.Next(ctx) {
			var user model.User
			err := rows.Decode(&user)
			if err != nil {
				log.Panic(err)
			}
			users[cnt] = user
			cnt += 1
		}
		if err := rows.Err(); err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
	}
}

func GetPotentialFriends(c *gin.Context) {
	userID := c.MustGet("userID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if userID != "" && err == nil {
		users := make([]model.User, pageSize)
		rows, err := database.GetPotentialFollowing(userID.(string), index, pageSize)
		if err != nil {
			log.Panic(err)
		}
		defer rows.Close(ctx)
		cnt := 0
		for rows.Next(ctx) {
			var user model.User
			err := rows.Decode(&user)
			if err != nil {
				log.Panic(err)
			}
			users[cnt] = user
			cnt += 1
		}
		if err := rows.Err(); err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
	}
}

func AddFriend(c *gin.Context) {
	userID := c.MustGet("userID")
	if userID != "" {
		var follower model.Follower
		if err := c.Bind(&follower); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message": "ok"})
		}
		follower.UserID = userID.(string)
		result, err := database.AddFollowing(follower)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusAccepted, "data": result.InsertedID, "message": "accepted",
		})
	}
}

func RemoveFriend(c *gin.Context) {
	userID := c.MustGet("userID")
	if userID != "" {
		var follower model.Follower
		if err := c.Bind(&follower); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message": "ok"})
		}
		follower.UserID = userID.(string)
		result, err := database.RemoveFollowing(follower)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusAccepted, "data": result.DeletedCount, "message": "accepted",
		})
	}
}
