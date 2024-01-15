package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type AppEnvironment string

const (
	PRODUCTION  AppEnvironment = "prod"
	STAGE       AppEnvironment = "stage"
	DEVELOPMENT AppEnvironment = "dev"
	LOCAL       AppEnvironment = "local"
	QA          AppEnvironment = "qa"
)

type (
	Config struct {
		Service      *Service
		Database     *Database
		Integrations *Integrations
	}

	Service struct {
		Port        string         `envconfig:"port" default:"8000"`
		Environment AppEnvironment `envconfig:"ENVIRONMENT" default:"local"`
	}

	Database struct {
		PostgreDSN string `envconfig:"POSTGRE_DSN" required:"true"`
	}

	Integrations struct {
		CentrifugoServerUrl     string `envconfig:"CENTRIFUGO_SERVER_URL" required:"true"`
		CentrifugoServerXApiKey string `envconfig:"CENTRIFUGO_SERVER_API_KEY" required:"true"`
	}
)

var (
	once   sync.Once
	config *Config
)

// GetConfig Загружает конфиг из .env файла и возвращает объект конфигурации
// В случае, если не передать параметр `envfiles`, берется `.env` файл из корня проекта
func GetConfig(envfiles ...string) (*Config, error) {
	var err error
	once.Do(func() {
		_ = godotenv.Load(envfiles...)

		var c Config
		err = envconfig.Process("", &c)
		if err != nil {
			err = fmt.Errorf("error parse config from env variables: %w\n", err)
			return
		}

		config = &c
	})

	return config, err
}
