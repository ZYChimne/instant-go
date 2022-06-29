package main

import (
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/router"
)

func main() {
	database.ConnectRedis()
	database.ConnectSQL()
	router.Create()
}
