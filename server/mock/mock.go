package mock

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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
	db.SetLogger(log.New(os.Stdout, "", 0))
	db.LogMode(true)

	models.DB = db
	models.Migrate()

	return err
}

// NewMockLogger creates a new mock logger
func NewMockLogger() {
	logger.New(logger.Mconsole, logger.Ldebug)
}
