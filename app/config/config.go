package config

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Read a YAML configuration and create a Config object out of it.
func LoadConfig(configPath string, log *logrus.Logger) (*Config, error) {
	viper.SetEnvPrefix("VEGA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.BindEnv("addr")
	_ = viper.BindEnv("tls_cert")
	_ = viper.BindEnv("tls_key")
	_ = viper.BindEnv("casbin_model")
	_ = viper.BindEnv("database.password")
	_ = viper.BindEnv("redis.password")
	_ = viper.BindEnv("session.secret")

	// viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "./config.yml")
	// because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("unable to find config file" + configPath)
		}
	}

	var cfg Config
	_ = viper.Unmarshal(&cfg)

	// set log level
	switch cfg.LogLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
		break
	case "info":
		log.SetLevel(logrus.InfoLevel)
		break
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "production" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else if env == "dev" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.JSONFormatter{})
		// The TextFormatter is default, you don't actually have to do this.
		// log.SetFormatter(&logrus.TextFormatter{})
	}
	log.WithFields(logrus.Fields{
		"config": cfg,
	}).Info("===> Vega is running in development mode. <===")

	return &cfg, nil
}
