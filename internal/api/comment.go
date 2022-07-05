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
	userID := c.MustGet("UserID")
	insID := c.Query("insID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if userID != "" && err == nil  {
		comments := make([]model.Comment, pageSize)
		rows, err := database.GetComments(insID, index, pageSize)
		if err != nil {
			log.Panic(err)
		}
		defer rows.Close(ctx)
		cnt := 0
		for rows.Next(ctx) {
			var comment model.Comment
			err := rows.Decode(&comment)
			if err != nil {
				log.Panic(err)
			}
			comments[cnt] = comment
			cnt += 1
		}
		if err := rows.Err(); err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": comments})
	}
}

func PostComment(c *gin.Context) {
	userID := c.MustGet("UserID")
	if userID !="" {
		var comment model.Comment
		if err := c.Bind(&comment); err != nil {
			log.Panic("Bind json failed ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error()})
		}
		comment.UserID = userID.(string)
		result, err := database.PostComment(comment)
		if err != nil {
			log.Panic("Post comment error ", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK, "data": result.InsertedID,
		})
	}
}