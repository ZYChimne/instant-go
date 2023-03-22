package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

const pageSize int64 = 10
const redisExpireTime time.Duration = 0 // 0 means no expire, ONLY FOR DEBUG

func Abort(c *gin.Context, err error, code int, message string) {
	if(err != nil) {
		log.Println(message, err.Error())
	}else{
		log.Println(message)
	}
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{"code": code, "message": message},
	)
}

