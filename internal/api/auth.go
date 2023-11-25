package api

import (
	"errors"
	"log"
	"net/http"
	"net/mail"
	database "zychimne/instant/internal/db"
	"zychimne/instant/tools"
	"zychimne/instant/pkg/model"
	"zychimne/instant/pkg/schema"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	var userSchema schema.LoginAccountRequest
	if err := c.Bind(&userSchema); err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LoginError))
		return
	}
	if len(userSchema.Email) > 0 {
		if _, err := mail.ParseAddress(userSchema.Email); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LoginError))
			return
		}
		var userModel model.User
		if err := database.LoginUserByEmail(userSchema.Email, &userModel); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(LoginError))
			return
		}
		if !util.CheckPasswordHash(userSchema.Password, userModel.Password) {
			c.AbortWithError(http.StatusUnauthorized, errors.New(LoginError))
			return
		}
		token := util.GenerateJwt(userModel.ID)
		c.JSON(http.StatusOK, gin.H{"data": token})
		return
	}
	if len(userSchema.Phone) > 0 {
		for _, char := range userSchema.Phone {
			if char < '0' || char > '9' {
				c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LoginError))
				return
			}
		}
		var userModel model.User
		if err := database.LoginUserByPhone(userSchema.Phone, &userModel); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(LoginError))
			return
		}
		if !util.CheckPasswordHash(userSchema.Password, userModel.Password) {
			c.AbortWithError(http.StatusUnauthorized, errors.New(LoginError))
			return
		}
		token := util.GenerateJwt(userModel.ID)
		c.JSON(http.StatusOK, gin.H{"data": token})
		return
	}
	if len(userSchema.Username) > 0 {
		var userModel model.User
		if err := database.LoginUserByUsername(userSchema.Username, &userModel); err != nil {
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError, errors.New(LoginError))
			return
		}
		if !util.CheckPasswordHash(userSchema.Password, userModel.Password) {
			c.AbortWithError(http.StatusUnauthorized, errors.New(LoginError))
			return
		}
		token := util.GenerateJwt(userModel.ID)
		c.JSON(http.StatusOK, gin.H{"data": token})
		return
	}
	c.AbortWithError(http.StatusUnprocessableEntity, errors.New(LoginError))
}
