package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	_ = c.MustGet("UserID").(uint)
	instantID, err := strconv.ParseUint(c.Query("instantID"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetCommentsError))
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetCommentsError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetCommentsError))
		return
	}
	var comments []model.Comment
	err = database.GetComments(uint(instantID), int(offset), int(limit), &comments)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetCommentsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func AddComment(c *gin.Context) {
	userID := c.MustGet("UserID").(uint)
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(AddCommentError))
		return
	}
	comment.UserID = userID
	err := database.AddComment(&comment)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(AddCommentError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": comment.ID})
}
