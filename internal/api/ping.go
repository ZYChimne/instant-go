package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	slp := c.Query("sleep")
	if len(slp) > 0 {
		slp, _ := strconv.Atoi(slp)
		time.Sleep(time.Duration(slp) * time.Millisecond)
	}
	c.String(http.StatusOK, "pong")
}
