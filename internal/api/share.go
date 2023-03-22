package api

import (
	"net/http"
	"strconv"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetShares(c *gin.Context) {
	_ = c.MustGet("UserID")
	errMsg := "Get shares error"
	insID := c.Query("insID")
	index, err := strconv.ParseInt(c.Query("index"), 10, 64)
	if err != nil {
		Abort(c, err, http.StatusBadRequest, errMsg)
		return
	}
	rows, err := database.GetShares(insID, index, pageSize)
	if err != nil {
		Abort(c, err, http.StatusBadRequest, errMsg)
		return
	}
	defer rows.Close(ctx)
	shares := []model.Share{}
	for rows.Next(ctx) {
		var sharing model.Share
		err := rows.Decode(&sharing)
		if err != nil {
			Abort(c, err, http.StatusBadRequest, errMsg)
			return
		}
		shares = append(shares, sharing)
	}
	if err := rows.Err(); err != nil {
		Abort(c, err, http.StatusBadRequest, errMsg)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": shares})
}

func PostSharingInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	errMsg := "Post sharing instants error"
	var share_sentence model.Share
	if err := c.Bind(&share_sentence); err != nil {
		Abort(c, err, http.StatusBadRequest, errMsg)
		return
	}
	share_sentence.UserID = userID.(string)
	result, err := database.PostShare(share_sentence)
	if err != nil {
		Abort(c, err, http.StatusBadRequest, errMsg)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": result.InsertedID,
	})
}
