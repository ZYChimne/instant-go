package api

import (
	"container/list"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

const pageSize int64 = 10
const errorListMaxSize int = 0x10000 // 2^17=65536

var redisExpireTime time.Duration = 0 // 0 means no expire, ONLY FOR DEBUG
const redisMaxMiss int = 0x10000      // 2^17=65536
const errorExpireTime int64 = 60      // unit: seconds

const (
	Warning        = 0
	UndefinedError = 1
	JsonError      = 2
	DatabaseError  = 3
	RedisError     = 4
	PasswordError  = 5
)

var UndefinedErrorList *list.List = list.New()
var JsonErrorList *list.List = list.New()
var DatabaseErrorList *list.List = list.New()
var RedisErrorList *list.List = list.New()
var PasswordErrorList *list.List = list.New()

func updateErrorList(cur int64, l *list.List) {
	l.PushBack(cur)
	for l.Len() > 0 && cur-l.Front().Value.(int64) > errorExpireTime {
		l.Remove(l.Front())
	}
	if l.Len() > errorListMaxSize {
		l.Remove(l.Front())
	}
}

func updateRedisExpireTime() {
	if DatabaseErrorList.Len() > 0 {
		redisExpireTime = time.Duration(DatabaseErrorList.Len()) * time.Second
	} else {
		redisExpireTime = 1 * time.Second
	}
}

func handleError(c *gin.Context, err error, code int, message string, errCode int) {
	if err != nil {
		log.Println(message, err.Error())
	} else {
		log.Println(message)
	}
	cur := time.Now().Unix()
	switch errCode {
	case Warning:
		{
			// DO NOTHING
		}
	case UndefinedError:
		{
			updateErrorList(cur, UndefinedErrorList)
		}
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": code, "message": message},
		)
	case JsonError:
		{
			updateErrorList(cur, JsonErrorList)
		}
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": code, "message": message},
		)
	case DatabaseError:
		{
			updateErrorList(cur, DatabaseErrorList)

		}
		updateRedisExpireTime()
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": code, "message": message},
		)
	case RedisError:
		{
			updateErrorList(cur, RedisErrorList)
			if RedisErrorList.Len() >= redisMaxMiss {
				redisExpireTime = -1
			} else {
				updateRedisExpireTime()
			}
		}
	case PasswordError:
		{
			updateErrorList(cur, PasswordErrorList)
		}
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": code, "message": message},
		)
	}
}
