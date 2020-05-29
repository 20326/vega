package db

import (
	"time"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // mysql
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite
)

type ()

// New db connect to a database
func NewDB(config Config) (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(config.Driver, config.DSN)
	if nil != err {
		// log.Fatal().Err(err).Str("driver", c.Driver).Str("dsn", c.DSN).Msg("open database failed")
		return nil, err
	}
	db.DB().SetMaxIdleConns(config.MaxOpenConns)                                    // default: 10
	db.DB().SetMaxOpenConns(config.MaxOpenConns)                                    // default: 50
	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute) // default: 5
	db.LogMode(config.LogMode)

	//cleanFunc := func() {
	//	err := db.Close()
	//	if nil != err {
	//		log.Fatal().Err(err).Str("driver", driver).Str("dsn", dsn).Msg("gorm close db failed")
	//	}
	//}
	// log.Debug().Msg("connected db")

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.TablePrefix + defaultTableName
	}

	// auto migrate database
	if config.AutoMigrate {
		err = AutoMigrateDB(db)
	}
	return db, err
}

func AutoMigrateDB(db *gorm.DB) error {
	// var err error
	print("auto migrate db")
	return db.AutoMigrate(
		new(model.Setting),
		new(model.Action),
		new(model.Resource),
		new(model.Permission),
		new(model.Role),
		new(model.User),
	).Error
}
