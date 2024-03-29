package cfgldr

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
	CorsConfig     CorsConfig
}

type DatabaseConfig struct {
	Url string
}

type AppConfig struct {
	Port   string
	Env    string
	ApiKey string
}

type CorsConfig struct {
	AllowOrigins string
}

func LoadConfig() (*Config, error) {
	if os.Getenv("CORS_ORIGINS") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	dbConfig := DatabaseConfig{
		Url: os.Getenv("DB_URL"),
	}

	appConfig := AppConfig{
		Port:   os.Getenv("APP_PORT"),
		Env:    os.Getenv("APP_ENV"),
		ApiKey: os.Getenv("APP_API_KEY"),
	}

	corsConfig := CorsConfig{
		AllowOrigins: os.Getenv("CORS_ORIGINS"),
	}

	return &Config{
		AppConfig:      appConfig,
		DatabaseConfig: dbConfig,
		CorsConfig:     corsConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
