package events

import "encoding/json"

// Join is a event for joining a channel.
type Join struct {
	Channel  string `json:"channel"`
	Password string `json:"password"`
}

// Part is a event for leaving a channel.
type Part struct {
	Channel string `json:"channel"`
}

// Msg is a event for sending a message
type Msg struct {
	From   string `json:"from,omitempty"`
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
	Target string `json:"target"`
	Msg    string `json:"msg"`
	Notice bool   `json:"notice"`
}

// Joined returns a join event message.
func Joined(channel string) string {
	event, err := json.Marshal(StatusMsgEvent{Type: "join",
		Status: "ok",
		Msg:    channel,
	})

	if err != nil {
		return InternalError(err.Error())
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
		return InternalError(err.Error())
	}

	return string(event)
}

// RcvedMsg returned a msg event.
func RcvedMsg(from, target, msg string, notice bool) string {
	event, err := json.Marshal(Msg{Type: "msg",
		From:   from,
		Status: "ok",
		Msg:    msg,
		Notice: notice,
		Target: target,
	})

	if err != nil {
		return InternalError(err.Error())
	}

	return string(event)
}
