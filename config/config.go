package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env     string `mapstructure:"env"`
	Address string `mapstructure:"address"`

	Database struct {
		Driver string `mapstructure:"driver"`
		Host   string `mapstructure:"host"`
		Port   int    `mapstructure:"port"`
		Name   string `mapstructure:"name"`

		Username string
		Password string
	} `mapstructure:"database"`
}

var AppConfig *Config

func Load() {
	// 1) Load .env (secrets)
	_ = godotenv.Load("local.env")

	// 2) Setup viper for yaml
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config.yaml: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// 3) Inject secrets from env into struct
	cfg.Database.Username = os.Getenv("DB_USERNAME")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")

	AppConfig = &cfg
}
