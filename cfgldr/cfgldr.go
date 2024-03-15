package cfgldr

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	AppConfig      AppConfig
	CorsConfig     CorsConfig
}

type DatabaseConfig struct {
	Url string `mapstructure:"URL"`
}

type AppConfig struct {
	Port   string `mapstructure:"PORT"`
	Env    string `mapstructure:"ENV"`
	ApiKey string `mapstructure:"API_KEY"`
}

type CorsConfig struct {
	AllowOrigins string `mapstructure:"ORIGINS"`
}

func LoadConfig() (*Config, error) {
	dbCfgLdr := viper.New()
	dbCfgLdr.SetEnvPrefix("DB")
	dbCfgLdr.AutomaticEnv()
	dbCfgLdr.AllowEmptyEnv(false)
	dbConfig := DatabaseConfig{}
	if err := dbCfgLdr.Unmarshal(&dbConfig); err != nil {
		return nil, err
	}

	appCfgLdr := viper.New()
	appCfgLdr.SetEnvPrefix("APP")
	appCfgLdr.AutomaticEnv()
	dbCfgLdr.AllowEmptyEnv(false)
	appConfig := AppConfig{}
	if err := appCfgLdr.Unmarshal(&appConfig); err != nil {
		return nil, err
	}

	corsConfigLdr := viper.New()
	corsConfigLdr.SetEnvPrefix("CORS")
	corsConfigLdr.AutomaticEnv()
	dbCfgLdr.AllowEmptyEnv(false)
	corsConfig := CorsConfig{}
	if err := corsConfigLdr.Unmarshal(&corsConfig); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseConfig: dbConfig,
		AppConfig:      appConfig,
		CorsConfig:     corsConfig,
	}, nil
}

func (ac *AppConfig) IsDevelopment() bool {
	return ac.Env == "development"
}
