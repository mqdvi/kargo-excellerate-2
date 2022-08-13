package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName      string `envconfig:"APP_NAME"`
	AppHost      string `envconfig:"APP_HOST"`
	SlugAppName  string `envconfig:"SLUG_APP_NAME"`
	AppQuote     string `envconfig:"APP_QUOTE"`
	Environment  string `envconfig:"ENVIRONMENT"`
	AppPort      string `envconfig:"APP_PORT"`
	PrintStatLog string `envconfig:"PRINT_STAT_LOG"`

	DBHost         string `envconfig:"DB_HOST"`
	DBPort         string `envconfig:"DB_PORT"`
	DBUser         string `envconfig:"DB_USER"`
	DBPassword     string `envconfig:"DB_PASSWORD"`
	DBName         string `envconfig:"DB_NAME"`
	DBConnection   string `envconfig:"DB_CONNECTION"`
	DBDialTimeout  int    `envconfig:"DB_DIAL_TIMEOUT"`
	DBReadTimeout  int    `envconfig:"DB_READ_TIMEOUT"`
	DBWriteTimeout int    `envconfig:"DB_WRITE_TIMEOUT"`
	DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`

	RedisHost     string `envconfig:"REDIS_HOST"`
	RedisPort     string `envconfig:"REDIS_PORT"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`

	BaseURL string `envconfig:"BASE_URL"`

	DefaultPaginationLimit uint64 `envconfig:"DEFAULT_PAGINATION_LIMIT"`
}

var once sync.Once
var instance Config

// GetConfig reads the configuration from .env file
func GetConfig() Config {
	once.Do(func() {
		_ = godotenv.Load()
		err := envconfig.Process("", &instance)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return instance
}
