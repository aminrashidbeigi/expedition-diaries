package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment     string
	Host            string
	Url             string
	Port            string
	Sitemaplocation string
	DBConfig        DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

var default_confs = map[string]string{
	"environment": "production",
	"host":        "0.0.0.0",
	"port":        "8080",
}

func LoadConfig(configFile string) (*Config, error) {
	for k, v := range default_confs {
		viper.SetDefault(k, v)
	}
	var cfg Config

	if configFile != "" {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
