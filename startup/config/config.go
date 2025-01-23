package config

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port int      `mapstructure:"port"`
	DB   DBConfig `mapstructure:"database"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func NewConfig() (*Config, error) {
	// First try to load from CONFIG_JSON environment variable
	if configJSON := os.Getenv("CONFIG_JSON"); configJSON != "" {
		var cfg Config
		err := json.Unmarshal([]byte(configJSON), &cfg)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
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
