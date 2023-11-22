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

// The Feed in an inbox with a message upper bound
// It does not retrieve messages beyond that upper bound (ex. Instagram)
// Instead, display trending post for the user
func GetFeed(c *gin.Context) {
	userID := c.MustGet("UserID")
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFeedError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFeedError))
		return
	}
	instants := []model.Feed{}
	err = database.GetFeed(userID.(uint), int(offset), int(limit), &instants)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetFeedError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": instants})
}
