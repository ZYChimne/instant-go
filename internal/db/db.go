package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:20001006@(127.0.0.1:3306)/instant?parseTime=true")
	if err != nil {
		log.Fatal("database error: ", err.Error())
	}
	return db
}
