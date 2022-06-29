package api

import (
	"log"
	"net/http"
	database "zychimne/instant/internal/db"
	"zychimne/instant/pkg/model"

	"github.com/gin-gonic/gin"
)

func GetFriends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": "", "message":"ok"})
}

func GetPotentialFriends(c *gin.Context) {
	var users []model.User
	rows, err := database.GetPotentialFriends(14)
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
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": err.Error(), "message":"ok"})
	}
	// Sort Before Insert
	if(friend.FirstID>friend.SecondID){
		temp:=friend.FirstID
		friend.FirstID=friend.SecondID
		friend.SecondID=temp
	}
	result, err := database.AddFriend(friend)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Post instant error ", err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusAccepted, "data": id, "message":"accepted",
	})
}

func RemoveFriend(c *gin.Context) {
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		log.Fatal("Bind json failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "data": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK, "data": "","message":"ok",
	})
}
