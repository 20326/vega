package middleware


import (
	"github.com/20326/vega/app/model"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

func SessionsMiddleware(sessions model.Session) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := sessions.Get(c)
		log.Error().Err(err).Msgf("session user mw: %+v", user)

		c.Next()
	}
}
