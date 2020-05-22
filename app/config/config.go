package config

import (
	"errors"
	"os"
	"strings"

	"github.com/phuslu/log"
	"github.com/spf13/viper"
)

// Read a YAML configuration and create a Config object out of it.
func LoadConfig(configPath string) (*Config, error) {
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

	var conf Config
	_ = viper.Unmarshal(&conf)

	// set log level
	switch conf.LogLevel {
	case "debug":
		log.DefaultLogger.SetLevel(log.DebugLevel)
		break
	case "info":
		log.DefaultLogger.SetLevel(log.InfoLevel)
		break

	case "warn":
		log.DefaultLogger.SetLevel(log.WarnLevel)
	}

	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		log.Info().Msg("===> Vega is running in development mode. <===")
	}
	log.Error().Str("ENVIRONMENT", env).Msgf("Vega config: %+v", &conf)

	return &conf, nil
}
