package settings

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	// mysql's connection string required by the gorm mysql driver
	MySQLDSN string `envconfig:"MYSQL_DSN" required:"true"`
	// the address the web server listens on
	ServerAddr string `envconfig:"SERVER_ADDR" default:":8080"`
}

// Load loads the env vars from the .env file, serializes them into a Settings struct
func Load() (*Settings, error) {
	// load env vars from the .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	configMap := &Settings{}

	// serialize the env vars into configMap
	err = envconfig.Process("", configMap)
	if err != nil {
		return nil, err
	}

	return configMap, nil
}
