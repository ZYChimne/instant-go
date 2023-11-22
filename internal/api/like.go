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

func Like(c *gin.Context) {
	userID := c.MustGet("UserID")
	var like model.InstantLike
	if err := c.Bind(&like); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LikeError))
		return
	}
	like.UserID = userID.(uint)
	err := database.LikeInstant(&like)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(LikeError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": like.ID})
}

func Unlike(c *gin.Context) {
	userID := c.MustGet("UserID")
	var like model.InstantLike
	if err := c.Bind(&like); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LikeError))
		return
	}
	like.UserID = userID.(uint)
	err := database.UnlikeInstant(&like)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(LikeError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": like.ID})
}

func GetLikes(c *gin.Context) {
	_ = c.MustGet("UserID")
	instantID, err := strconv.ParseUint(c.Query("instantID"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetLikesError))
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetLikesError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetLikesError))
		return
	}
	var likes []model.InstantLike
	err = database.GetLikes(uint(instantID), int(offset), int(limit), &likes)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetLikesError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": likes})
}
