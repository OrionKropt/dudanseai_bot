package configs

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Config struct {
	TelegramApiKey string
	LogLevel       string
	ProxyUrl       string
	DbBot          DbConfig
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ReadConfig() {
	cfg.TelegramApiKey = os.Getenv("TELEGRAM_API_KEY")
	cfg.LogLevel = os.Getenv("LOG_LEVEL")
	cfg.ProxyUrl = os.Getenv("PROXY_URL")
	cfg.DbBot.User = os.Getenv("DB_USER")
	cfg.DbBot.Password = os.Getenv("DB_PASSWORD")
	cfg.DbBot.Host = os.Getenv("DB_HOST")
	cfg.DbBot.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	cfg.DbBot.Name = os.Getenv("DB_NAME")
}
