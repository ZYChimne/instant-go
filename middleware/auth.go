package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"zychimne/instant/internal/util"

	"github.com/gin-gonic/gin"
)

const AuthError = "token is invalid"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithError(http.StatusUnauthorized, errors.New(AuthError))
			return
		}
		userID, err := util.VerifyJwt(token[7:])
		if err != nil {
			log.Println(err.Error())
			c.AbortWithError(http.StatusUnauthorized, errors.New(AuthError))
			return
		}
		c.Set("UserID", userID)
	}
}
