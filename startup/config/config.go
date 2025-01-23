package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
	// Try to load from environment variables first
	if os.Getenv("DB_HOST") != "" {
		// Debug: Print environment variables
		fmt.Printf("DB_HOST: %s\n", os.Getenv("DB_HOST"))
		fmt.Printf("DB_PORT: %s\n", os.Getenv("DB_PORT"))
		fmt.Printf("DB_USER: %s\n", os.Getenv("DB_USER"))
		fmt.Printf("DB_NAME: %s\n", os.Getenv("DB_NAME"))

		dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, fmt.Errorf("invalid DB_PORT: %v", err)
		}

		configJSON := map[string]interface{}{
			"port": 3000,
			"database": map[string]interface{}{
				"host":     os.Getenv("DB_HOST"),
				"port":     dbPort,
				"user":     os.Getenv("DB_USER"),
				"password": os.Getenv("DB_PASSWORD"),
				"dbname":   os.Getenv("DB_NAME"),
				"sslmode":  "require",
			},
		}

		jsonBytes, _ := json.Marshal(configJSON)
		var cfg Config
		if err := json.Unmarshal(jsonBytes, &cfg); err != nil {
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
