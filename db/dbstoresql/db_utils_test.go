package dbstoresql_test

import (
	"hungknow.com/blockchain/db/dbconfig"
	"hungknow.com/blockchain/db/dbstoresql"
	"hungknow.com/blockchain/errors"
)

func PrepareDbStore() (*dbstoresql.DBSQLStore, *errors.AppError) {
	// Get the integration test
	config := &dbconfig.Config{
		Host:         "localhost",
		Port:         5432,
		Username:     "postgres",
	}
	return dbstoresql.NewDBSQLStore(config)
}
