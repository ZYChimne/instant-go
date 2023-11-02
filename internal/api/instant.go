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

func GetInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	_targetID := c.Query("userID")
	if len(_targetID) == 0 {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetInstantsError))
		return
	}
	targetID, err := strconv.ParseUint(_targetID, 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetInstantsError))
		return
	}
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetInstantsError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetInstantsError))
		return
	}
	var instants []model.Instant
	err = database.GetInstants(userID.(uint), uint(targetID), int(offset), int(limit), &instants)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetInstantsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": instants})
}

func AddInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(AddInstantError))
		return
	}
	instant.UserID = userID.(uint)
	err := database.AddInstant(&instant)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(AddInstantError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": instant.ID})
}

func UpdateInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(AddInstantError))
		return
	}
	instant.UserID = userID.(uint)
	err := database.UpdateInstant(&instant)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(AddInstantError))
		return
	}
	c.Status(http.StatusOK)
}

func DeleteInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	_instantID := c.Query("instantID")
	if len(_instantID) == 0 {
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(DeleteInstantError))
		return
	}
	instantID, err := strconv.ParseUint(_instantID, 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(DeleteInstantError))
		return
	}
	var instant model.Instant
	instant.ID = uint(instantID)
	instant.UserID = userID.(uint)
	err = database.DeleteInstant(&instant)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(DeleteInstantError))
		return
	}
	c.Status(http.StatusOK)
}
