package middleware

import (
	"fmt"
	"log"
	"strconv"

	"github.com/20326/vega/app/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func NewSessionsStore(config *config.Config) sessions.Store {
	store, err := redis.NewStoreWithDB(
		config.Redis.MaxConn,
		"tcp",
		fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		config.Redis.Password,
		strconv.Itoa(config.Redis.DBIndex),
		[]byte(config.Session.Secret))

	if nil != err {
		log.Fatalf("create session redis store failed, err: %s", err)
	}

	store.Options(sessions.Options{
		// Domain:   conf.Session.Domain,
		Path:     "/",
		MaxAge:   int(config.Session.Expiration),
		Secure:   false, //conf.Session.Secure,//conf.TLSCert != "", // TODO
		HttpOnly: true,
	})

	_ = redis.SetKeyPrefix(store, config.Session.KeyPrefix+":")
	return store
}
