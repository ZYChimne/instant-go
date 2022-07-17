package api

import (
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	_ = c.MustGet("UserID")
	errMsg := "Get comments error"
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
	rows, err := database.GetComments(insID, index, pageSize)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	defer rows.Close(ctx)
	comments := []model.Comment{}
	for rows.Next(ctx) {
		var comment model.Comment
		err := rows.Decode(&comment)
		if err != nil {
			log.Println("Database Decode error ", err.Error())
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "message": errMsg},
			)
			return
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": comments})
}

func PostComment(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Post comment error"
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Println("Bind json failed ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	comment.UserID = userID.(string)
	result, err := database.PostComment(comment)
	if err != nil {
		log.Println("Database error ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "message": errMsg},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": result.InsertedID,
	})
}
