package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func GetFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFollowingsError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFollowingsError))
		return
	}
	followings := []model.JointFollowing{}
	err = database.GetFollowings(userID.(uint), int(offset), int(limit), &followings)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetFollowingsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": followings})
}

func GetFollowers(c *gin.Context) {
	userID := c.MustGet("UserID")
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFollowersError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFollowersError))
		return
	}
	followers := []model.JointFollowing{}
	err = database.GetFollowers(userID.(uint), int(offset), int(limit), &followers)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetFollowersError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": followers})
}

func GetPotentialFollowings(c *gin.Context) {
	userID := c.MustGet("UserID")
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetPotentialFollowingsError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetPotentialFollowingsError))
		return
	}
	users := []model.User{}
	err = database.GetPotentialFollowings(userID.(uint), int(offset), int(limit), &users)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetPotentialFollowingsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Follow(c *gin.Context) {
	userID := c.MustGet("UserID")
	var followingSchema schema.UpsertFollowingRequest
	if err := c.Bind(&followingSchema); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(FollowError))
		return
	}
	following := model.Following{
		UserID:   userID.(uint),
		TargetID: followingSchema.TargetID,
	}
	err := database.AddFollowing(&following)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(FollowError))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": following.ID})
}

func Unfollow(c *gin.Context) {
	userID := c.MustGet("UserID")
	var followingSchema schema.UpsertFollowingRequest
	if err := c.Bind(&followingSchema); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(UnfollowError))
		return
	}
	following := model.Following{
		UserID:   userID.(uint),
		TargetID: followingSchema.TargetID,
	}
	err := database.RemoveFollowing(&following)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(UnfollowError))
		return
	}
	c.Status(http.StatusOK)
}

func GetFriends(c *gin.Context) {
	userID := c.MustGet("UserID")
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFriendsError))
		return
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(GetFriendsError))
		return
	}
	friends := []model.JointFollowing{}
	err = database.GetFollowings(userID.(uint), int(offset), int(limit), &friends)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New(GetFriendsError))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": friends})
}
