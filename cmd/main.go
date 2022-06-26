package main

import (
	database "zychimne/instant/internal/db"
	"zychimne/instant/internal/router"
)

func main() {
	database.ConnectRedis()
	router.Create()
}
