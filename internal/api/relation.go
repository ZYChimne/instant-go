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
	errMsg := "Get followings error"
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		log.Println("Parse index err ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	followings := []model.Following{}
	rows, err := database.GetFollowings(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var following model.Following
		err := rows.Decode(&following)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
			return
		}
		followings = append(followings, following)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": followings})
}

func GetFollowers(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followers error"
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		log.Println("Parse index err ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	followers := []model.Following{}
	rows, err := database.GetFollowers(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var follower model.Following
		err := rows.Decode(&follower)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
			return
		}
		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": followers})
}

func GetPotentialFollowings(c *gin.Context) {
	// userID := c.MustGet("UserID")
	// errMsg := "Get potential following error"
	// index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	// if err != nil {
	// 	log.Println("Parse index error ", err.Error())
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
	// 	return
	// }
	// followingOIDs := []primitive.ObjectID{}
	// rows, err := database.GetFollowings(userID.(string), index, math.MaxInt64)
	// if err != nil {
	// 	log.Println("Database error ", err.Error())
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
	// 	return
	// }
	// defer rows.Close(ctx)
	// for rows.Next(ctx) {
	// 	var following model.Following
	// 	err := rows.Decode(&following)
	// 	if err != nil {
	// 		log.Println("Database error ", err.Error())
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
	// 		return
	// 	}
	// 	followingOID,err :=primitive.ObjectIDFromHex(following.FollowingID)
	// 	if err != nil {
	// 		log.Println("Database error ", err.Error())
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
	// 		return
	// 	}
	// 	followingOIDs = append(followingOIDs, followingOID)
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Println("Database error ", err.Error())
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": errMsg})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": followings})
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
		err := database.AddFollowing(following)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusAccepted, "message": "accepted",
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
		err := database.RemoveFollowing(following)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusAccepted, "message": "accepted",
		})
	}
}
