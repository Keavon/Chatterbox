package models

import (
	"github.com/jinzhu/gorm"
	// PQ loaded to use postgres driver.
	_ "github.com/lib/pq"

	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

// Global DB connection
var DB gorm.DB

// New initalizes the database
func New(con string) error {
	var err error
	DB, err = gorm.Open("postgres", con)
	DB.SetLogger(logger.Debug)
	DB.LogMode(true)
	return err
}
