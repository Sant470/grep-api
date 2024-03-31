package config

import (
	"log"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

type Config struct {
	AWS AWSCredential
}

type AWSCredential struct {
	AccessKeyID     string
	SecretAccessKey string
}

func GetAppConfig(filename, path string) *Config {
	return loadConfig(filename, path)
}

func InitRouters() *chi.Mux {
	r := chi.NewRouter()
	// setup cors here ...
	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
	)
	return r
}

func loadConfig(filename, path string) *Config {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("fatal: error reading config file")
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal("fatal: error reading config variable")
	}
	return &conf
}
