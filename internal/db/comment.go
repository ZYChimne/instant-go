package database

import (
	"database/sql"
	"log"
	"time"
	"zychimne/instant/pkg/model"
)

func GetComments(insID int) (*sql.Rows, error) {
	query := "SELECT comment_id, create_time, update_time, user_id, content FROM comments WHERE ins_id = ?"
	return db.Query(query, insID)
}

func PostComment(comment model.Comment) (sql.Result, error) {
	query := `INSERT INTO comments (create_time, update_time, ins_id, user_id, content) VALUES (?, ?, ?, ?, ?)`
	return db.Exec(query, time.Now(), time.Now(), comment.InsID, comment.UserID, comment.Content)
}

func LikeComment() {
	log.Println("Not implemented")
}
