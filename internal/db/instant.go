package database

import (
	"database/sql"
	"time"
	"zychimne/instant/pkg/model"
)

func GetInstants(userID any, index string) (*sql.Rows, error){
	query := "SELECT ins_id, create_time, update_time, content FROM instants WHERE user_id = ? ORDER BY ins_id DESC LIMIT ?, 10"
	return db.Query(query, userID, index)
}

func PostInstant(instant model.Instant) (sql.Result, error){
	query := `INSERT INTO instants (user_id, create_time, update_time, content) VALUES (?, ?, ?, ?)`
	return db.Exec(query, instant.UserID, time.Now(), time.Now(), instant.Content)
}

func UpdateInstant(instant model.Instant) (sql.Result, error){
	query := `UPDATE instants SET update_time = ?, content = ? WHERE ins_id = ?`
	return db.Exec(query, time.Now(), instant.Content, instant.InsID)
}

func LikeInstant(like model.Like) (sql.Result, error){
	query := `INSERT INTO likes (create_time, update_time, ins_id, user_id, attitude) VALUES (?, ?, ?, ?, ?)`
	return db.Exec(query, time.Now(), time.Now(), like.InsID, like.UserID, like.Attitude)
}

func ShareInstant(instant model.Instant)(sql.Result, error){
		query := `INSERT INTO instants (user_id, create_time, update_time, content, ref_origin_id) VALUES (?, ?, ?, ?, ?)`
		return db.Exec(query, instant.UserID, time.Now(), time.Now(), instant.Content, instant.RefOriginId)
}