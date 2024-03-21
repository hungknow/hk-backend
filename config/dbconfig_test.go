package config_test

import (
	"testing"

	"github.com/knadh/koanf/v2"
	"github.com/stretchr/testify/require"
	"hungknow.com/blockchain/config"
	"hungknow.com/blockchain/db/dbconfig"
)

func TestDBCon(t *testing.T) {
	k := koanf.New(".")
	expectedDBConfig := &dbconfig.DBConfig{
		Host: "localhost",
		Port: 8888,
		Username: "postgres",
		Password: "password",
		DatabaseName: "test",
	}
	k.Set("database", *expectedDBConfig)

	dbConfig, err := config.LoadDBConfig(k)
	require.NoError(t, err)
	require.Equal(t, expectedDBConfig, dbConfig)

	
	// Test the username from environment variable
	expectedDBConfig1 := &*expectedDBConfig
	expectedDBConfig1.Username = "test1"
	k.Set("DATABASE_USERNAME", expectedDBConfig1.Username)
	dbConfig, err = config.LoadDBConfig(k)
	require.NoError(t, err)
	require.Equal(t, expectedDBConfig1, dbConfig)

	expectedDBConfig2 := &*expectedDBConfig1
	expectedDBConfig2.Password = "password2"
	k.Set("DATABASE_PASSWORD", "password2") 
	dbConfig, err = config.LoadDBConfig(k)
	require.NoError(t, err)
	require.Equal(t, expectedDBConfig2, dbConfig)
}