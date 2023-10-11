package api

import (
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func GetFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followings error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	followings := []model.Following{}
	err = database.GetFollowings(userID.(uint), int(offset), int(limit), &followings)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": followings})
}

func GetJointFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followings error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	followings := []model.JointFollowing{}
	err = database.GetJointFollowings(userID.(uint), int(offset), int(limit), &followings)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": followings})
}

func GetFollowers(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followings error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	followers := []model.Following{}
	err = database.GetFollowers(userID.(uint), int(offset), int(limit), &followers)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": followers})
}

func GetJointFollowers(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followings error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	followers := []model.JointFollowing{}
	err = database.GetJointFollowers(userID.(uint), int(offset), int(limit), &followers)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": followers})
}

func GetPotentialFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get potential following error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, UndefinedError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	users := []model.User{}
	err = database.GetPotentialFollowings(userID.(uint), int(offset), int(limit), &users)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": users})
}

func AddFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Add following error"
	var followingSchema schema.UpdateFollowingRequest
	if err := c.Bind(&followingSchema); err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	following:=model.Following{
		UserID: userID.(uint),
		TargetID: followingSchema.TargetID,
	}	
	err := database.AddFollowing(&following)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})
}

func RemoveFollowing(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Remove following error"
	var followingSchema schema.UpdateFollowingRequest
	if err := c.Bind(&followingSchema); err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	following:=model.Following{
		UserID: userID.(uint),
		TargetID: followingSchema.TargetID,
	}	
	err := database.RemoveFollowing(&following)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

func GetFriends(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followings error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	friends := []model.Following{}
	err = database.GetFollowings(userID.(uint), int(offset), int(limit), &friends)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": friends})
}


func GetJointFriends(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Get followings error"
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		handleError(c, err, errMsg, ParameterError)
		return
	}
	friends := []model.JointFollowing{}
	err = database.GetJointFollowings(userID.(uint), int(offset), int(limit), &friends)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": friends})
}