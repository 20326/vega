package middleware

import (
	"github.com/20326/vega/app/service"

	"github.com/gin-gonic/gin"
)

func ServiceMiddleware(srv *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		srv.WithContext(c)
		c.Next()
	}
}
