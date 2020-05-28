package redis

type Config struct {
	Host     string
	Port     int64
	Password string
	DBIndex  int
	MaxConn  int
}
