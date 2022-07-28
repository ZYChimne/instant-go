package api

import (
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetSharings(c *gin.Context) {
	_ = c.MustGet("UserID")
	errMsg := "Get sharings error"
	insID := c.Query("insID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		log.Println("Parse index error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	rows, err := database.GetSharings(insID, index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	sharings := []model.Sharing{}
	for rows.Next(ctx) {
		var sharing model.Sharing
		err := rows.Decode(&sharing)
		if err != nil {
			log.Println("Database Decode error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		sharings = append(sharings, sharing)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": sharings})
}

func PostSharingInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Post sharing instants error"
	var share_sentence model.Sharing
	if err := c.Bind(&share_sentence); err != nil {
		log.Println("Bind json error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	share_sentence.UserID = userID.(string)
	result, err := database.PostSharing(share_sentence)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": result.InsertedID,
	})
}
