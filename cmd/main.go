package main

import (
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/router"

	"zychimne/instant/config"
)

func main() {
	config.LoadConfig()
	database.ConnectPostgres()
	database.ConnectRedis()
	router.Create()
}
