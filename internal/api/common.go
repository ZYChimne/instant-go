package api

import (
	"container/list"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/util/hotkey"
	"zychimne/instant/internal/util/lru"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

const pageSize int64 = 10

const errorListMaxSize int = 0x10000 // 2^17=65536

var redisExpireTime time.Duration = 0 // 0 means no expire, ONLY FOR DEBUG
const redisMaxMiss int = 0x10000      // 2^17=65536
const errorExpireTime int64 = 60      // unit: seconds
const redisLikeInstantKeysSet string = "like_instant_keys"

const (
	Warning        = 0
	UndefinedError = 1
	JsonError      = 2
	DatabaseError  = 3
	RedisError     = 4
	PasswordError  = 5
	BindError      = 6
)

var UndefinedErrorList *list.List = list.New()
var JsonErrorList *list.List = list.New()
var DatabaseErrorList *list.List = list.New()
var RedisErrorList *list.List = list.New()
var PasswordErrorList *list.List = list.New()
var BindErrorList *list.List = list.New()

func getFromCache(c *gin.Context, key string, localCache *hotkey.HotKeyCache) bool {
	item, ok := localCache.Get(key)
	if ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": item})
		return true
	} else if getFromRedis(c, key) {
		return true
	}
	return false
}

func getFromRedis(c *gin.Context, key string) bool {
	res, err := database.RedisClient.Get(ctx, key).Result()
	if err != nil {
		handleError(nil, err, 0, "Get from redis error", RedisError)
		return false
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": res})
	return true
}

func putInRedis(key string, item interface{}) {
	var res string
	if reflect.TypeOf(item).Kind() == reflect.String {
		res = item.(string)
	} else {
		bytes, err := json.Marshal(item)
		if err != nil {
			handleError(nil, err, 0, "Marshal error", JsonError)
		}
		res = string(bytes)
	}
	if redisExpireTime >= 0 {
		err := database.RedisClient.Set(ctx, key, res, redisExpireTime).Err()
		if err != nil {
			handleError(nil, err, 0, "Put in redis error", RedisError)
		}
	}
}

func putInCache(key string, item any, localCache *hotkey.HotKeyCache) {
	if !localCache.Add(key, item) {
		putInRedis(key, item)
	}
}

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

func handleError(c *gin.Context, err error, httpCode int, message string, errCode int) {
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
			gin.H{"code": httpCode, "message": message},
		)
	case JsonError:
		{
			updateErrorList(cur, JsonErrorList)
		}
	case DatabaseError:
		{
			updateErrorList(cur, DatabaseErrorList)

		}
		updateRedisExpireTime()
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"code": httpCode, "message": message},
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
			gin.H{"code": httpCode, "message": message},
		)
	case BindError:
		{
			updateErrorList(cur, BindErrorList)
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"code": httpCode, "message": message},
			)
		}
	}
}

func onLikeInstantRedis(key string, incr int64) {
	pipe := database.RedisClient.TxPipeline()
	pipe.SAdd(ctx, redisLikeInstantKeysSet, key)
	pipe.IncrBy(ctx, key, incr)
	_, err := pipe.Exec(ctx)
	if err != nil {
		handleError(nil, err, 0, "onLikeInstantRedis", RedisError)
	}
}

func onLikeInstantEvicted(key lru.Key, value interface{}) {
	onLikeInstantRedis(key.(string), value.(int64))
}
