package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		c.JSON(-1, gin.H{"message": c.Errors.Last().Error()})
		for _, err := range c.Errors {
			log.Println(err.Meta, err.Type, err.Error())
		}
	}
}
