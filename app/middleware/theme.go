package middleware

import (
	"github.com/gin-gonic/gin"
)

func ThemeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
