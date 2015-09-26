package events

import (
	"encoding/json"
	"fmt"
)

// Connected returns a connection event
func Connected(server string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "connection",
		Status: "ok",
		Msg: fmt.Sprintf("connected to %s",
			server),
	})

	if err != nil {
		return internalError(err.Error())
	}

	return string(event)
}

// WaitForConnection returns a connection in progress message
func WaitForConnection(time float64) string {
	event, err := json.Marshal(StatusMsgDurationEvent{Type: "connection",
		Status:   "working",
		Msg:      "waiting for connection",
		Duration: time,
	})

	if err != nil {
		return internalError(err.Error())
	}

	return string(event)
}
