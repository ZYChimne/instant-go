package api

import (
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get instants error"
	indexStr := c.Query("index")
	index, err := strconv.ParseInt(indexStr, 0, 64)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	rows, err := database.GetInstants(userID.(string), index, pageSize)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	defer rows.Close(ctx)
	instants := []model.Instant{}
	for rows.Next(ctx) {
		var instant model.Instant
		err := rows.Decode(&instant)
		if err != nil {
			handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
			return
		}
		instants = append(instants, instant)
	}
	if err := rows.Err(); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instants})
}

func GetInstantsByUserID(c *gin.Context) {
	userID := c.Query("userID")
	indexStr := c.Query("index")
	errMsg := "Get instants error"
	index, err := strconv.ParseInt(indexStr, 0, 64)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	rows, err := database.GetInstantsByUserID(userID, index, pageSize)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	defer rows.Close(ctx)
	instants := []model.Instant{}
	for rows.Next(ctx) {
		var instant model.Instant
		err := rows.Decode(&instant)
		if err != nil {
			handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
			return
		}
		instants = append(instants, instant)
	}
	if err := rows.Err(); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instants})
}

func PostInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Post instant error"
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	instant.UserID = userID.(string)
	err := database.PostInstant(instant)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
	})
}

func UpdateInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Update instant error"
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	instant.UserID = userID.(string)
	result, err := database.UpdateInstant(instant)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	if result.ModifiedCount == 0 {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusCreated,
	})
}

func LikeInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Like instant error"
	var like model.Like
	if err := c.Bind(&like); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	like.UserID = userID.(string)
	err := database.LikeInstant(like)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
	})
}

func ShareInstant(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Share instant error"
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	instant.UserID = userID.(string)
	_, err := database.ShareInstant(instant)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func GetLikesUserInfo(c *gin.Context) {
	_ = c.MustGet("UserID")
	errMsg := "Get like description error"
	insID := c.Query("insID")
	indexStr := c.Query("index")
	index, err := strconv.ParseInt(indexStr, 0, 64)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, BindError)
		return
	}
	rows, err := database.GetLikesUserInfo(insID, index, pageSize)
	if err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	defer rows.Close(ctx)
	userInfos := []model.SimpleUser{}
	for rows.Next(ctx) {
		var userInfo model.SimpleUser
		err := rows.Decode(&userInfo)
		if err != nil {
			handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
			return
		}
		userInfos = append(userInfos, userInfo)
	}
	if err := rows.Err(); err != nil {
		handleError(c, err, http.StatusBadRequest, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": userInfos,
	})
}
