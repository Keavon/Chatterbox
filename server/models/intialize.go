package models

import (
	"github.com/jinzhu/gorm"
	// PQ loaded to use postgres driver.
	_ "github.com/lib/pq"
)

// Global DB connection
var DB gorm.DB

// New initalizes the database
func New(con string) error {
	var err error
	DB, err = gorm.Open("postgres", con)
	return err
}
