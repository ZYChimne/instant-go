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

func GetShares(c *gin.Context) {
	_ = c.MustGet("UserID")
	instantID, err := strconv.ParseUint(c.Query("instantID"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetSharesError))
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetSharesError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetSharesError))
		return
	}
	var instants []model.Instant
	err = database.GetShares(uint(instantID), int(offset), int(limit), &instants)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetSharesError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": instants})
}

func ShareInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(ShareInstantError))
		return
	}
	instant.UserID = userID.(uint)
	err := database.ShareInstant(&instant)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(ShareInstantError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": instant.ID})
}
