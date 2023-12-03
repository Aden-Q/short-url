package setting

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Setting struct {
	// the timeout for the routing pipeline to process a request, measured in seconds
	RequestTimeout int `envconfig:"REQUEST_TIMEOUT" default:"60"`
	// mysql's connection string required by the gorm mysql driver
	MySQLDSN string `envconfig:"MYSQL_DSN" required:"true"`
	// address the web server listens on
	ServerAddr string `envconfig:"SERVER_ADDR" default:":8080"`
	// address of the redis server
	RedisAddr string `envconfig:"REDIS_ADDR" default:":6379"`
}

// Load loads the env vars from the .env file, serializes them into a Settings struct
func Load() (*Setting, error) {
	// load env vars from the .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	configMap := &Setting{}

	// serialize the env vars into configMap
	err = envconfig.Process("", configMap)
	if err != nil {
		return nil, err
	}

	return configMap, nil
}
