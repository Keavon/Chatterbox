package models

import (
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"

	"github.com/wayn3h0/go-uuid"
	"github.com/wayn3h0/go-uuid/random"
)

// GenerateUUID generates a uuid.
func generateUUID() (string, error) {
	nUUID, err := random.New()

	if err != nil {
		logger.Error.Print(err)
		return "", err
	}

	return nUUID.Format(uuid.StyleWithoutDash), nil
}
