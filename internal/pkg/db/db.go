package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBEngine struct {
	*gorm.DB
}

func NewDBEngine() (*DBEngine, error) {
	dns := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(
		// use settings from docker-compose.yml
		// override the default env var if inside docker container
		mysql.Open(dns),
		&gorm.Config{},
	)

	return &DBEngine{db}, err
}
