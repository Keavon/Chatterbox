package validate

import "regexp"

// ValidationMsg is a the message format of validations.
type ValidationMsg struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

const (
	// UniqueMsg is the error if something isn't unique.
	UniqueMsg = "must be unique"
	// NilMsg is error if something is nil.
	NilMsg = "can't be nil"
	// NotEmailMsg is returned when an string isn't an email.
	NotEmailMsg = "not a valid email"
)

var (
	emailRegex = regexp.MustCompile(`^[A-z0-9._%+-]+@[A-z0-9.-]+\.[A-z]{2,}$`)
)

// IsEmail validates that a string is a email.
func IsEmail(field, email string) []ValidationMsg {
	if emailRegex.MatchString(email) {
		return []ValidationMsg{}
	}

	return []ValidationMsg{ValidationMsg{Field: field, Msg: NotEmailMsg}}
}

// NotNil validates a sting isn't empty
func NotNil(field, contents string) []ValidationMsg {
	if contents == "" {
		return []ValidationMsg{ValidationMsg{Field: field, Msg: NilMsg}}
	}

	return []ValidationMsg{}
}

// ValidationMsgsToJSON converts a ValidationMsg array to a json error message.
func ValidationMsgsToJSON(msgs []ValidationMsg) interface{} {
	return map[string]interface{}{
		"errors": msgs,
	}
}
