package models

import (
	"encoding/json"
	"regexp"

	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

// ValidationMsg is a the message format of validations.
type ValidationMsg struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

const (
	uniqueMsg   = "must be unique"
	nilMsg      = "can't be nil"
	notEmailMsg = "not a valid email"
)

var (
	emailRegex = regexp.MustCompile(`^[A-z0-9._%+-]+@[A-z0-9.-]+\.[A-z]{2,}$`)
)

func isEmail(field, email string) []ValidationMsg {
	if emailRegex.MatchString(email) {
		return []ValidationMsg{}
	}

	return []ValidationMsg{ValidationMsg{Field: field, Msg: notEmailMsg}}
}

func notNil(field, contents string) []ValidationMsg {
	if contents == "" {
		return []ValidationMsg{ValidationMsg{Field: field, Msg: nilMsg}}
	}

	return []ValidationMsg{}
}

// ValidationToJSON converts a ValidationMsg array to a json error message.
func ValidationToJSON(msgs []ValidationMsg) ([]byte, error) {
	errMsg := map[string]interface{}{
		"error": msgs,
	}

	out, err := json.Marshal(errMsg)

	if err != nil {
		logger.Error.Print(err)
		return nil, err
	}

	return out, nil
}
