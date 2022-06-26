package api

import (
	"log"
	"net/http"
	"time"
	"zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetFriends(c *gin.Context) {
	var comments []model.Comment
	query := "SELECT commentid, create_time, update_time, userid, content FROM comments WHERE insid = ?"
	db := database.ConnectDatabase()
	rows, err := db.Query(query, 1)
	// db.Close()
	// defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.CommentID, &comment.CreateTime, &comment.UpdateTime, &comment.UserID, &comment.Content)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": comments})
}

func GetPotentialFriends(c *gin.Context) {
	var users []model.User
	// query := "SELECT firstid, secondid FROM friends WHERE firstid = ? OR secondid = ?"
	query := "SELECT userid, avatar, username FROM accounts WHERE userid != ?"
	db := database.ConnectDatabase()
	rows, err := db.Query(query, 14)
	// db.Close()
	// defer rows.Close()
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Avatar, &user.Username)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": users})
}

func AddFriend(c *gin.Context) {
	var friend model.Friend
	if err := c.Bind(&friend); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	// Sort Before Insert
	if(friend.FirstID>friend.SecondID){
		temp:=friend.FirstID
		friend.FirstID=friend.SecondID
		friend.SecondID=temp
	}
	query := `INSERT INTO friends (create_time, update_time, firstid, secondid) VALUES (?, ?, ?, ?)`
	db := database.ConnectDatabase()
	result, err := db.Exec(query, time.Now(), time.Now(), friend.FirstID, friend.SecondID)
	db.Close()
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200, "id": id,
	})
}

func RemoveFriend(c *gin.Context) {
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO comments (create_time, update_time, ins_id, user_id, content) VALUES (?, ?, ?, ?, ?)`
	db := database.ConnectDatabase()
	result, err := db.Exec(query, time.Now(), time.Now(), comment.InsID, comment.UserID, comment.Content)
	db.Close()
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200, "id": id,
	})
}
