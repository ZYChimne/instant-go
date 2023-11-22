package api

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func Receive(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", "Hello world!")
		time.Sleep(1 * time.Second)
		return true
	})
}

func Send(c *gin.Context) {

}
