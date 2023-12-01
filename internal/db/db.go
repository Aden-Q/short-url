package db

import (
	"github.com/Aden-Q/short-url/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type Config struct {
	MySQLDSN string
	gorm.Config
}

type Engine interface {
	// First finds the first record that matches the given conditions, order by primary key
	First(dest interface{}, conds ...interface{}) error
	// Create creates a new record
	Create(value interface{}) error
	// Update updates the value of a column in a table
	Update(table interface{}, column string, value interface{}) error
}

type engine struct {
	*gorm.DB
}

// NewEngine creates a new database engine and establishes a connection to the database
func NewEngine(config Config) (Engine, error) {
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

	return &engine{db}, err
}

func (e *engine) First(dest interface{}, conds ...interface{}) error {
	return e.DB.First(dest, conds...).Error
}

func (e *engine) Create(value interface{}) error {
	return e.DB.Create(value).Error
}

// Update updates the value of a column in a table
// table will be used to build the condition for the update
// column is the column to be updated
// value is the new value
func (e *engine) Update(table interface{}, column string, value interface{}) error {
	return e.DB.Model(table).Update(column, value).Error
}
