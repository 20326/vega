package db

type Config struct {
	Driver          string
	DSN             string
	TablePrefix     string
	AutoMigrate     bool
	LogMode         bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

// New db connect to a database
func NewConfig(driver, dsn, tablePrefix string, autoMigrate, logMode bool, maxIdelConns, maxOpenConns, connMaxLifetime int) Config {
	return Config{
		Driver:          driver,
		DSN:             dsn,
		TablePrefix:     tablePrefix,
		AutoMigrate:     autoMigrate,
		LogMode:         logMode,
		MaxIdleConns:    maxIdelConns,
		MaxOpenConns:    maxOpenConns,
		ConnMaxLifetime: connMaxLifetime,
	}
}
