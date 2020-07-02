package redis

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	DefaultPoolName = ""
)

type ClientPool struct {
	mutex   *sync.Mutex
	options *redis.Options
	pools   map[string]*redis.Client
}

func New(config *Config) *ClientPool {
	options := NewRedisOptions(config)
	// get default client
	client := GetClient(DefaultPoolName)
	if nil == client {
		logrus.Fatalf("new redis client failed")
	}

	return &ClientPool{
		mutex:   &sync.Mutex{},
		options: options,
		pools: map[string]*redis.Client{
			DefaultPoolName: client,
		},
	}
}

// NewClient net the client of redis.
func (srv *ClientPool) NewClient() (*redis.Client, error) {
	var client *redis.Client
	var err error

	client = redis.NewClient(srv.options)
	if nil == client {
		err = errors.New("new redis client failed")
	}

	_, err = client.Ping().Result()
	if err != nil {
		logrus.Fatalf("ping redis failed" + err.Error())
		return nil, err
	}

	logrus.WithError(err).WithField("client", client.String()).Printf("ping: " + pong)
	return client, err
}

func NewRedisOptions(config *Config) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DBIndex,
	}
}

func (srv *ClientPool) GetClient(key string) *redis.Client {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	var err error
	if _, ok := srv.pools[key]; !ok {
		srv.pools[key], err = srv.NewClient()
		if nil != err {
			logrus.WithError(err).Printf("new redis client failed")
			return nil
		}
	}

	return srv.pools[key]
}
