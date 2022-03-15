package apicomment

import (
	"log"
	"net/http"
	"time"
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var comments []model.Comment
	query := "SELECT commentid, create_time, update_time, userid, content FROM comments WHERE insid = ?"
	db := sql.ConnectDatabase()
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

func PostComment(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO comments (create_time, update_time, ins_id, user_id, content) VALUES (?, ?, ?, ?, ?)`
	db := sql.ConnectDatabase()
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

func LikeComment(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO comments (create_time, update_time, ins_id, user_id, content) VALUES (?, ?, ?, ?, ?)`
	db := sql.ConnectDatabase()
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

func ShareComment(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO comments (create_time, update_time, ins_id, user_id, content) VALUES (?, ?, ?, ?, ?)`
	db := sql.ConnectDatabase()
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
