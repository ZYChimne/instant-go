package database

import (
	"database/sql"
	"log"
	"time"
	"zychimne/instant/pkg/model"
)

func GetFriends() {
	log.Println("Not implemented")
}

func GetPotentialFriends(userID int) (*sql.Rows, error) {
	query := "SELECT user_id, avatar, username FROM accounts WHERE user_id != ?"
	return db.Query(query, userID)
}

func AddFriend(friend model.Friend) (sql.Result, error) {
	query := `INSERT INTO friends (create_time, update_time, first_id, second_id) VALUES (?, ?, ?, ?)`
	return db.Exec(query, time.Now(), time.Now(), friend.FirstID, friend.SecondID)
}

func RemoveFriend(friend model.Friend) {
	log.Println("Not implemented")
}
