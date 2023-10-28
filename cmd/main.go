package main

import (
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/router"

	"zychimne/instant/config"
)

func main() {
	config.LoadConfig("config/dev.yml")
	database.ConnectPostgres()
	database.ConnectRedis()
	r := router.Create(true)
	r.Run(":" + config.Conf.Instant.Port)
}
