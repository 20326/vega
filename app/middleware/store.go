package middleware

import (
	"github.com/20326/vega/app/service"

	"github.com/gin-gonic/gin"

)

func ServiceMiddleware(svr *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		svr.WithContext(c)
		c.Next()
	}
}
