package database

import (
	"strings"

	"github.com/redis/go-redis/v9"

	"zychimne/instant/config"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     strings.Join([]string{config.Conf.Database.Redis.Host, config.Conf.Database.Redis.Port}, ":"),
		Password: config.Conf.Database.Redis.Password,
		DB:       config.Conf.Database.Redis.Database,
	})
}
