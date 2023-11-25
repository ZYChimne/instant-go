package api

import (
	"io"
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

func EventPing(c *gin.Context) {
	slp := c.Query("sleep")
	if len(slp) == 0 {
		slp = "1000"
	}
	_slp, _ := strconv.Atoi(slp)
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("ping", "pong")
		time.Sleep(time.Duration(_slp) * time.Millisecond)
		return true
	})
}