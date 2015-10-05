package parser

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/chatterbox-irc/chatterbox/ircc/events"
	"github.com/chatterbox-irc/chatterbox/ircc/irc"
	"github.com/chatterbox-irc/chatterbox/pkg/validate"
)

// Parse parsses the input string and runs the command
// Returns -1 if successful, 0 to signal graceful exit, >0 to signify critical exit
func Parse(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := make(map[string]interface{})

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	if _, ok := cmd["type"].(string); !ok {
		fmt.Fprint(w, events.JSONError("type field must exist and be a string"))
		return -1 // Not a critical error.
	}

	switch cmd["type"].(string) {
	case "exit":
		return exit(ircc, w)
	case "join":
		return join(ircc, w, input)
	case "part":
		return part(ircc, w, input)
	case "msg":
		return msg(ircc, w, input)
	default:
		fmt.Fprint(w, events.JSONError(fmt.Sprintf("unknown type %s", cmd["type"].(string))))
		return -1 // Not a critical error.
	}
}

func exit(ircc *irc.IRC, w io.Writer) int {
	ircc.Disconnect()
	return 0
}

func join(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Join{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("channel", cmd.Channel)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("join", e))
		return -1 // Not a critical error.
	}

	ircc.Join(cmd.Channel, cmd.Password)
	return -1
}

func part(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Part{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("channel", cmd.Channel)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("part", e))
		return -1 // Not a critical error.
	}

	ircc.Part(cmd.Channel)
	return -1
}

// TODO: Add CTCP support for newlines
// TODO: check length of message and break message into multiple lines
func msg(ircc *irc.IRC, w io.Writer, input string) int {
	cmd := events.Msg{}

	if err := json.Unmarshal([]byte(input), &cmd); err != nil {
		fmt.Fprint(w, events.JSONError(err.Error()))
		return -1 // Not a critical error.
	}

	e := []validate.ValidationMsg{}
	e = append(e, validate.NotNil("target", cmd.Target)...)
	e = append(e, validate.NotNil("msg", cmd.Msg)...)

	if len(e) > 0 {
		fmt.Fprint(w, events.ValidationError("join", e))
		return -1 // Not a critical error.
	}

	ircc.Msg(cmd.Target, cmd.Msg, cmd.Notice)
	return -1
}
