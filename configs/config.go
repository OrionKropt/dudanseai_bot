package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	TelegramApiKey string `json:"telegram_api_key"`
	LogLevel       string `json:"log_level"`
	configPath     string
}

func NewConfig() *Config {

	return &Config{
		LogLevel:   "debug",
		configPath: os.Getenv("CONFIG_PATH"),
	}
}

func (cfg *Config) ReadConfig() error {
	data, err := os.ReadFile(cfg.configPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return err
	}

	return nil
}
