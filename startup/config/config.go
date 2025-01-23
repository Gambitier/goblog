package config

import (
	"encoding/json"
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
		dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		appPort, _ := strconv.Atoi(os.Getenv("PORT"))
		configJSON := map[string]interface{}{
			"port": appPort,
			"database": map[string]interface{}{
				"host":     os.Getenv("DB_HOST"),
				"port":     dbPort,
				"user":     os.Getenv("DB_USER"),
				"password": os.Getenv("DB_PASSWORD"),
				"dbname":   os.Getenv("DB_NAME"),
				"sslmode":  "require",
			},
		}

		// Convert to JSON and use existing unmarshal logic
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
