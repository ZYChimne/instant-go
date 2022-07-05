package api

import (
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if userID != "" && err == nil {
		followings := []model.Following{}
		rows, err := database.GetPotentialFollowing(userID.(string), index, pageSize)
		if err != nil {
			log.Panic(err)
		}
		defer rows.Close(ctx)
		for rows.Next(ctx) {
			var following model.Following
			err := rows.Decode(&following)
			if err != nil {
				log.Panic(err)
			}
			followings=append(followings, following)
		}
		if err := rows.Err(); err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": followings})
	}
}

func GetPotentialFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if userID != "" && err == nil {
		followings := []model.Following{}
		rows, err := database.GetPotentialFollowing(userID.(string), index, pageSize)
		if err != nil {
			log.Panic(err)
		}
		defer rows.Close(ctx)
		for rows.Next(ctx) {
			var following model.Following
			err := rows.Decode(&following)
			if err != nil {
				log.Panic(err)
			}
			followings=append(followings, following)
		}
		if err := rows.Err(); err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": followings})
	}
}

func AddFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	if userID != "" {
		var following model.Following
		if err := c.Bind(&following); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message": "ok"})
		}
		following.UserID = userID.(string)
		result, err := database.AddFollowing(following)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusAccepted, "data": result.InsertedID, "message": "accepted",
		})
	}
}

func RemoveFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	if userID != "" {
		var following model.Following
		if err := c.Bind(&following); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message": "ok"})
		}
		following.UserID = userID.(string)
		result, err := database.RemoveFollowing(following)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusAccepted, "data": result.DeletedCount, "message": "accepted",
		})
	}
}
