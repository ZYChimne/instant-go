package api

import (
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	index, err := strconv.ParseInt(c.Query("index"), 0, 64)
	if userID != nil && userID != "" && err == nil {
		instants := []model.Instant{}
		rows, err := database.GetInstants(userID.(string), index, pageSize)
		if err != nil {
			log.Panic(err.Error())
		}
		defer rows.Close(ctx)
		for rows.Next(ctx) {
			var instant model.Instant
			err := rows.Decode(&instant)
			if err != nil {
				log.Fatal(err)
			}
			instants = append(instants, instant)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instants, "message": "ok"})
	}
}

func PostInstant(c *gin.Context) {
	userID := c.MustGet("userID")
	if userID != "" {
		var instant model.Instant
		if err := c.Bind(&instant); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message": "error"})
		}
		instant.UserID = userID.(string)
		result, err := database.PostInstant(instant)
		if err != nil {
			log.Fatal("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, "data": result.InsertedID,
		})
	}
}

func UpdateInstant(c *gin.Context) {
	userID := c.MustGet("userID")
	if userID != "" {
		var instant model.Instant
		if err := c.Bind(&instant); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error()})
		}
		instant.UserID = userID.(string)
		result, err := database.UpdateInstant(instant)
		if err != nil {
			log.Fatal("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, "data": result.UpsertedID,
		})
	}
}

func LikeInstant(c *gin.Context) {
	userID := c.MustGet("userID")
	if userID != "" {
		var like model.Like
		if err := c.Bind(&like); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error()})
		}
		like.UserID = userID.(string)
		result, err := database.LikeInstant(like)
		if err != nil {
			log.Panic("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, "data": result.UpsertedID,
		})
	}
}

func ShareInstant(c *gin.Context) {
	userID := c.MustGet("userID")
	if userID != "" {
		var instant model.Instant
		if err := c.Bind(&instant); err != nil {
			log.Fatal("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
		}
		instant.UserID = userID.(string)
		result, err := database.ShareInstant(instant)
		if err != nil {
			log.Fatal("Post instant error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, "data": result.InsertedID,
		})
	}
}
