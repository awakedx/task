package config

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DbURI          string
	HTTPAddr       string
	MigrationsPath string
	ReadTimeOut    time.Duration
	WriteTimeOut   time.Duration
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			slog.Error("failed to load .env file", "error", err.Error())
		}
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../../config/")
		viper.AddConfigPath("./config")

		err = viper.ReadInConfig()
		if err != nil {
			slog.Error("failed to read .yaml file", "error", err.Error())
			panic(err)
		}
		config = Config{
			DbURI: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME")),
			HTTPAddr:       ":" + viper.GetString("server.port"),
			ReadTimeOut:    viper.GetDuration("server.readTimeOut"),
			WriteTimeOut:   viper.GetDuration("server.writeTimeOut"),
			MigrationsPath: viper.GetString("database.migrationsPath"),
		}
		slog.Info("Config was successfully read")
	})
	return &config
}
