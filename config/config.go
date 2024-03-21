package config

import (
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"hungknow.com/blockchain/db/dbconfig"
	"hungknow.com/blockchain/errors"
)

const (
	Local string = "local"
)

type AppConfig struct {
	DB *dbconfig.Config
}

func GetConfig() (*AppConfig, *errors.AppError) {
	k := koanf.New(".")
	dotenvParser := dotenv.Parser()
	yamlParser := yaml.Parser()
	err := k.Load(file.Provider(".env"), dotenvParser)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppConfigError, "Failed to load .env: %v", err)
	}
	err = k.Load(file.Provider("config/local.yml"), yamlParser)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppConfigError, "Failed to load config yaml file: %v", err)
	}

	dbConfig, appErr := GetDBConfig(k)
	if appErr != nil {
		return nil, appErr
	}

	allConfig := &AppConfig{
		DB: dbConfig,
	}

	return allConfig, nil
}
