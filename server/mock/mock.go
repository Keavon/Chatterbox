package mock

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	// SQLite loaded as driver.
	_ "github.com/mattn/go-sqlite3"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

// NewMockDB creates a new mock database
func NewMockDB() error {
	var err error
	db, err := gorm.Open("sqlite3", fmt.Sprintf("/tmp/test-%s-%s.db", time.Now().String(), string(rand.Int())))

	models.DB = db
	models.Migrate()

	return err
}

// NewMockLogger creates a new mock logger
func NewMockLogger() {
	format := log.Ldate | log.Ltime | log.Lshortfile
	logger.Debug = log.New(ioutil.Discard, "", format)
	logger.Info = log.New(ioutil.Discard, "", format)
	logger.Warn = log.New(ioutil.Discard, "", format)
	logger.Error = log.New(ioutil.Discard, "", format)
}
