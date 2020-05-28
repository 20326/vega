package middleware

import (
	"fmt"
	"strconv"

	"github.com/20326/vega/app/config"
	// "github.com/20326/vega/app/model"
	// "github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/phuslu/log"
)

func NewSessionsStore(config *config.Config) sessions.Store {
	store, err := redis.NewStoreWithDB(
		config.Redis.MaxConn,
		"tcp",
		fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		config.Redis.Password,
		strconv.Itoa(config.Redis.DBIndex),
		[]byte(config.Session.Secret))

	log.Info().Str("Secret", config.Session.Secret).Msgf("session Secret")

	if nil != err {
		log.Fatal().Err(err).Msg("create session redis store failed")
	}

	log.Info().Msgf("sessionStore: %+v", store)
	store.Options(sessions.Options{
		// Domain:   conf.Session.Domain,
		Path:     "/",
		MaxAge:   int(config.Session.Expiration),
		Secure:   false, //conf.Session.Secure,//conf.TLSCert != "", // TODO
		HttpOnly: true,
	})

	_ = redis.SetKeyPrefix(store, "session:")
	return store
}

//
//func SessionsMiddleware(sessions model.Session) gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//		user, err := sessions.Get(c)
//		log.Error().Err(err).Msgf("session user mw: %+v", user)
//
//		c.Next()
//	}
//}
