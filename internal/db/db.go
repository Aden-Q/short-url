package db

import (
	"github.com/Aden-Q/short-url/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	MySQLDSN string
	gorm.Config
}

type Engine struct {
	*gorm.DB
}

// NewEngine creates a new database engine and establishes a connection to the database
func NewEngine(config Config) (*Engine, error) {
	db, err := gorm.Open(
		// use settings from docker-compose.yml
		// override the default env var if inside docker container
		mysql.Open(config.MySQLDSN),
		&config.Config,
	)
	if err != nil {
		return nil, err
	}

	// run an initial migration
	if err = db.AutoMigrate(&model.URL{}); err != nil {
		// TODO: add a log
		return nil, err
	}

	// set an initial value for the auto increment key
	db.Exec("ALTER TABLE urls AUTO_INCREMENT = 100000")

	return &Engine{db}, err
}
