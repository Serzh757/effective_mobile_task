package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string

	InfoBackend string
}

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Error reading config file")
		return nil
	}
	config := &Config{
		Host:        viper.GetString("DB_HOST"),
		Port:        viper.GetInt("DB_PORT"),
		User:        viper.GetString("DB_USER"),
		Password:    viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
		InfoBackend: viper.GetString("INFO_BACKEND_URL"),
	}
	return config
}
