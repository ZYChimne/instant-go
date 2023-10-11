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
		handleError(c, err, errMsg, ParameterError)
		return
	}
	rows, err := database.GetShares(insID, index, pageSize)
	if err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	defer rows.Close(ctx)
	shares := []model.Share{}
	for rows.Next(ctx) {
		var sharing model.Share
		err := rows.Decode(&sharing)
		if err != nil {
			handleError(c, err, errMsg, DatabaseError)
			return
		}
		shares = append(shares, sharing)
	}
	if err := rows.Err(); err != nil {
		handleError(c, err, errMsg, DatabaseError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": shares})
}

