package apiinstant

import (
	"log"
	"net/http"
	"time"
	"zychimne/instant/internal/db"
	"zychimne/instant/internal/util"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetInstants(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 304})
		return
	}
	var instants []model.Instant
	query := "SELECT insid, create_time, update_time, content FROM instants WHERE userid = ?"
	db := sql.ConnectDatabase()
	rows, err := db.Query(query, 14)
	db.Close()
	defer rows.Close()
	for rows.Next() {
		var instant model.Instant
		err := rows.Scan(&instant.InsID, &instant.CreateTime, &instant.UpdateTime, &instant.Content)
		if err != nil {
			log.Fatal(err)
		}
		instants = append(instants, instant)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": instants})
}

func PostInstant(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO instants (userid, create_time, update_time, content) VALUES (?, ?, ?, ?)`
	db := sql.ConnectDatabase()
	result, err := db.Exec(query, instant.UserID, time.Now(), time.Now(), instant.Content)
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

func UpdateInstant(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `UPDATE instants SET update_time = ?, content = ? WHERE insid = ?`
	db := sql.ConnectDatabase()
	result, err := db.Exec(query, time.Now(), instant.Content, instant.InsID)
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

func LikeInstant(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var like model.Like
	if err := c.Bind(&like); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO likes (create_time, update_time, insid, userid, attitude) VALUES (?, ?, ?, ?, ?)`
	db := sql.ConnectDatabase()
	result, err := db.Exec(query, time.Now(), time.Now(), like.InsID, like.UserID, like.Attitude)
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

func ShareInstant(c *gin.Context) {
	token := c.GetHeader("Authentication")
	if err := utilauth.VerifyJwt(token); err != nil {
		log.Fatal("jwt error", err.Error())
	}
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO instants (userid, create_time, update_time, content, ref_origin_id) VALUES (?, ?, ?, ?, ?)`
	db := sql.ConnectDatabase()
	result, err := db.Exec(query, instant.UserID, time.Now(), time.Now(), instant.Content, instant.RefOriginId)
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
