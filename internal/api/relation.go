package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	followings := []model.Following{}
	rows, err := database.GetFollowings(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var following model.Following
		err := rows.Decode(&following)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		followings = append(followings, following)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
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
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	followers := []model.Following{}
	rows, err := database.GetFollowers(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var follower model.Following
		err := rows.Decode(&follower)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": followers})
}

func GetPotentialFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get potential following error"
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		log.Println("Parse index error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	users := []model.User{}
	rows, err := database.GetPotentialFollowings(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var user model.User
		err := rows.Decode(&user)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
}

func GetAllUsers(c *gin.Context) {
	index, err := strconv.ParseInt(c.Query("index"), 0, 64)
	errMsg := "Get all users error"
	if err != nil {
		log.Println("Parse index error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	users := []model.User{}
	rows, err := database.GetAllUsers(index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var user model.User
		err := rows.Decode(&user)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
}

func AddFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Add following error"
	var following model.Following
	if err := c.Bind(&following); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	following.UserID = userID.(string)
	err := database.AddFollowing(following)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	err = database.RedisClient.Del(ctx, strings.Join([]string{"recent", userID.(string)}, ":"), strings.Join([]string{"recent", following.FollowingID}, ":")).
		Err()
	if err != nil {
		log.Println("Redis error ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusAccepted, "message": "accepted",
	})
}

func RemoveFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Remove following error"
	var following model.Following
	if err := c.Bind(&following); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	following.UserID = userID.(string)
	err := database.RemoveFollowing(following)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	err = database.RedisClient.Del(ctx, strings.Join([]string{"recent", userID.(string)}, ":"), strings.Join([]string{"recent", following.FollowingID}, ":")).
		Err()
	if err != nil {
		log.Println("Redis error ", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusAccepted, "message": "accepted",
	})
}

func GetFriends(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get friends error"
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		log.Println("Parse index error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	key := strings.Join([]string{"recent", userID.(string)}, ":")
	users := []model.SimpleUser{}
	if index == 0 {
		if cache, err := database.RedisClient.Get(ctx, key).Result(); err != nil {
			log.Println(errMsg, " ", err.Error())
		} else {
			err := json.Unmarshal([]byte(cache), &users)
			if err != nil {
				log.Println("Unmarshal error ", err.Error())
			} else {
				c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
				return
			}
		}
	}
	rows, err := database.GetFriends(userID.(string), index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	for rows.Next(ctx) {
		var user model.SimpleUser
		err := rows.Decode(&user)
		if err != nil {
			log.Println("Database error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	if index == 0 {
		cache, err := json.Marshal(users)
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
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
}
