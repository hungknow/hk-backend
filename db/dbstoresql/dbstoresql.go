package dbstoresql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"hungknow.com/blockchain/db/dbconfig"
	"hungknow.com/blockchain/db/dbstore"
	"hungknow.com/blockchain/errors"
)

type DBSQLStore struct {
	db *sql.DB
}

func NewDBSQLStore(config *dbconfig.Config) (*DBSQLStore, *errors.AppError) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DatabaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.NewAppErrorf(errors.AppInternalError, "%v", err)
	}

	return &DBSQLStore{
		db: db,
	}, nil
}

func (o *DBSQLStore) Ping() error {
	return o.db.Ping()
}

func (o *DBSQLStore) Close() {
	o.db.Close()
}

func (o *DBSQLStore) ForexCandles() dbstore.ForexCandles {
	return o
}
