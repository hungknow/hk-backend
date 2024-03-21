package config

import (
	"github.com/knadh/koanf/v2"
	"hungknow.com/blockchain/db/dbconfig"
	"hungknow.com/blockchain/errors"
)

func GetDBConfig(k *koanf.Koanf) (*dbconfig.Config, *errors.AppError) {
	dbConfig := &dbconfig.Config{}
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
