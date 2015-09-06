package models

import (
	"fmt"
	"regexp"
)

var (
	uniqueErrTemplate = "%s must be unique"
	emailRegex        = regexp.MustCompile(`^[A-z0-9._%+-]+@[A-z0-9.-]+\.[A-z]{2,}$`)
)

func isEmail(email string) []string {
	if emailRegex.MatchString(email) {
		return []string{}
	}

	return []string{fmt.Sprintf("%s is not a valid email", email)}
}

func notNil(field, contents string) []string {
	if contents == "" {
		return []string{fmt.Sprintf("%s field can't be nil", field)}
	}

	return []string{}
}
