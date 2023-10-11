package database

import (
	"strings"

	"github.com/redis/go-redis/v9"

	"zychimne/instant/config"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     strings.Join([]string{config.Conf.Redis.Host, config.Conf.Redis.Port}, ":"),
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Database,
	})
}
