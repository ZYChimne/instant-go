package api

import (
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
		handleError(c, err, errMsg, ParameterError)
		return
	}
	rows, err := database.GetComments(insID, index, pageSize)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	defer rows.Close(ctx)
	comments := []model.Comment{}
	for rows.Next(ctx) {
		var comment model.Comment
		err := rows.Decode(&comment)
		if err != nil {
			handleError(c, err, errMsg, DatabaseError)
			return
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": comments})
}

func PostComment(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Post comment error"
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	comment.UserID = userID.(string)
	result, err := database.PostComment(comment) // TODO: use Redis
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": result.InsertedID,
	})
}
