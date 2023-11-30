package db

import (
	"os"

	"github.com/Aden-Q/short-url/internal/model"
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

	// run an initial migration
	db.AutoMigrate(&model.URL{})

	// set an initial value for the auto increment key
	db.Exec("ALTER TABLE urls AUTO_INCREMENT = 100000")

	return &DBEngine{db}, err
}
