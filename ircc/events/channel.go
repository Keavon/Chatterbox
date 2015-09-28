package events

import "encoding/json"

// Joined returns a join event message.
func Joined(channel string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "join",
		Status: "ok",
		Msg:    channel,
	})

	if err != nil {
		return internalError(err.Error())
	}

	return string(event)
}

// Parted returns a parted event message.
func Parted(channel string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "part",
		Status: "ok",
		Msg:    channel,
	})

	if err != nil {
		return internalError(err.Error())
	}

	return string(event)
}
