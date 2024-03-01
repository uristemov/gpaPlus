package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/uristemov/repeatPro/internal/applicator"
	"github.com/uristemov/repeatPro/internal/config"
	"github.com/uristemov/repeatPro/pkg/logger"
)

func main() {

	logger := logger.New()
	defer logger.Sync()

	cfg, err := loadConfig("config")
	if err != nil {
		logger.Panicf("failed to init config error: %v", err)
	}

	applicator.Run(logger, cfg)
}

func loadConfig(path string) (config *config.Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}

	if err = godotenv.Load(); err != nil {
		return config, fmt.Errorf("No .env file found err: %w", err)
	}

	return config, nil
}
