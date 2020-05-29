package config

import (
	"errors"
	"github.com/sirupsen/logrus"
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

	switch cfg.LogFormatter {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
		break
	case "text":
		log.SetFormatter(&logrus.TextFormatter{})
		break
	default:
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	log.WithFields(logrus.Fields{
		"config": cfg,
	}).Info("===> Vega is running in development mode. <===")

	return &cfg, nil
}
