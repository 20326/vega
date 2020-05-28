package config

// Config object extracted from YAML configuration file.

// DatabaseConfig represents the configuration of a mysql/sqlite3/postgres database
type DatabaseConfig struct {
	Driver          string
	DSN             string
	TablePrefix     string
	AutoMigrate     bool
	LogMode         bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

// RedisConfig represents the configuration related to redis cache & session store.
type RedisConfig struct {
	Host     string
	Port     int64
	Password string
	DBIndex  int
	MaxConn  int
}

// AdmissionConfig represents the configuration of casbin
type AdmissionConfig struct {
	CasbinModel string
	TablePrefix string
	LogMode     bool
}

type SessionConfig struct {
	Name       string
	Secret     string
	Expiration int64 // Expiration in seconds
	Inactivity int64 // Inactivity in seconds
	Domain     string
	Secure     bool
}

type Config struct {
	Addr        string
	PidFile     string
	TLSCert     string
	TLSKey      string
	LogLevel    string
	LogFilePath string
	Maintenance bool
	Database    DatabaseConfig
	Admission   AdmissionConfig
	Redis       RedisConfig
	Session     SessionConfig
}
