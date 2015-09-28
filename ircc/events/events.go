package events

// StatusMsgEvent is a event that includes a status and a message.
type StatusMsgEvent struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// StatusMsgDurationEvent is a event that includes a status, message, and duration.
type StatusMsgDurationEvent struct {
	Type     string  `json:"type"`
	Status   string  `json:"status"`
	Msg      string  `json:"msg"`
	Duration float64 `json:"duration"`
}

// StatusErrorsEvent is a event that includes a status and errors.
type StatusErrorsEvent struct {
	Type   string      `json:"type"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

// Join is a event for joining a channel.
type Join struct {
	Type     string `json:"type"`
	Channel  string `json:"channel"`
	Password string `json:"password"`
}

// Part is a event for leaving a channel.
type Part struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
}
