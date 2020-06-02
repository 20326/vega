package middleware

import (
	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"

	"net/http"
	"strings"
)

var ignoredPerms = map[string]bool{}

func AdmitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := strings.Split(c.Request.URL.RequestURI(), "?")[0]
		method := c.Request.Method

		srv := service.FromContext(c)
		log := srv.GetLogger()

		// TODO match /path/:id
		log.Warn("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\nAdmit: %s %s\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", method, path)
		if _, ok := ignoredPerms[path]; ok {
			c.Next()
			return
		}

		session := &model.SessionData{}
		user, _ := session.Get(c, srv.Users)

		allowed, _ := srv.Admissions.Admit(c, user, path, method)
		log.Warn("\n", method, path, allowed)
		if !allowed {
			log.Warn("No permission for %s %s", method, path)
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "err.Err403",
			})
			c.Abort()
			return
		} else {
			log.Info("permission check ok, %s %s", method, path)
		}
		c.Next()
	}
}
