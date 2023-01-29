package config

import (
	"elprup/internal/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Environment Env
	DbConfig    DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Fatal(err)
	}
	return &config
}
