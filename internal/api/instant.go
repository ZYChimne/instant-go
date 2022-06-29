package api

import (
	"log"
	"net/http"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetInstants(c *gin.Context) {
	userID := c.MustGet("UserID")
	index := c.Query("index")
	if userID != nil && userID != -1 && index != "" {
		instants := []model.Instant{}
		rows, err := database.GetInstants(userID, index)
		if err != nil {
			log.Fatal(err.Error())
		}
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
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message": "error"})
	}
	result, err := database.PostInstant(instant);
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
	result, err := database.UpdateInstant(instant)
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
	result, err := database.LikeInstant(like)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "id": id,
	})
}

func ShareInstant(c *gin.Context) {
	var instant model.Instant
	if err := c.Bind(&instant); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	result, err := database.ShareInstant(instant)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": id,
	})
}
