package config

import (
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/v2"
	"hungknow.com/blockchain/db/dbconfig"
	"hungknow.com/blockchain/errors"
)

var (
	// k            = koanf.New(".")
	dotenvParser = dotenv.Parser()
	yamlParser   = yaml.Parser()
	// parser = json.Parser()
)

const (
	Local string = "local"
)

func LoadDBConfig(k *koanf.Koanf) (*dbconfig.DBConfig, error) {
	dbConfig := &dbconfig.DBConfig{}
	err := k.Unmarshal("database", &dbConfig)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppInternalError, "Failed to load database config: %v", err)
	}

	databaseUserName := k.String("DATABASE_USERNAME")
	if databaseUserName != "" {
		dbConfig.Username = databaseUserName
	}

	databasePassword := k.String("DATABASE_PASSWORD")
	if databasePassword != "" {
		dbConfig.Password = databasePassword
	}

	return dbConfig, nil
}
