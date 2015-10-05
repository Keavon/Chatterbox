package events

import (
	"encoding/json"
	"fmt"

	"github.com/chatterbox-irc/chatterbox/pkg/validate"
)

// ConnectionError returns a connection error string.
func ConnectionError(msg string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "connection", Status: "error", Msg: msg})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

// JSONError returns a json parsing error string.
func JSONError(msg string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "json", Status: "error", Msg: msg})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

// ValidationError returns a validation error string.
func ValidationError(msgType string, msgs []validate.ValidationMsg) string {
	event, err := json.Marshal(StatusErrorsEvent{Type: msgType, Status: "failed", Errors: msgs})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}

// InternalError is user if there is a internal error that can't be recovered,
// so it can't throw a error.
func InternalError(msg string) string {
	return fmt.Sprintf(`{"type":"error","msg":"%s"}`, msg)
}
