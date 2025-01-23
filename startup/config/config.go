package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	DatabaseURL string `mapstructure:"database_url"`
}

func NewConfig() (*Config, error) {
	// Try to load from environment variables first
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return &Config{
			Port:        3000,
			DatabaseURL: dbURL,
		}, nil
	}

	// Fallback to file-based config for local development
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
