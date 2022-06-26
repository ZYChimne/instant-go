package api

import (
	"log"
	"net/http"
	"time"
	"zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	index := c.Query("index")
	if userID != nil && userID != -1 && index != "" {
		instants := []model.Instant{}
		query := "SELECT ins_id, create_time, update_time, content FROM instants WHERE user_id = ? LIMIT ?, 10"
		db := database.ConnectDatabase()
		rows, err := db.Query(query, userID, index)
		if err != nil {
			log.Fatal(err.Error())
		}
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
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": instants, "message": "ok"})
	}
}

func PostInstant(c *gin.Context) {
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO instants (user_id, create_time, update_time, content) VALUES (?, ?, ?, ?)`
	db := database.ConnectDatabase()
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
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `UPDATE instants SET update_time = ?, content = ? WHERE ins_id = ?`
	db := database.ConnectDatabase()
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
	var like model.Like
	if err := c.Bind(&like); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO likes (create_time, update_time, insid, userid, attitude) VALUES (?, ?, ?, ?, ?)`
	db := database.ConnectDatabase()
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
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	query := `INSERT INTO instants (userid, create_time, update_time, content, ref_origin_id) VALUES (?, ?, ?, ?, ?)`
	db := database.ConnectDatabase()
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
