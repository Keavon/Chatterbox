package auth

import (
	"testing"

	"github.com/chatterbox-irc/chatterbox/server/client"
)

type loginErrorCase struct {
	email    string
	password string
}

func TestLogin(t *testing.T) {
	api, server := NewAuthTest(t)
	defer server.Close()

	email, password := "t@t.com", "abc123"
	if _, err := api.Register(email, password); err != nil {
		t.Error(err)
	}

	testCases := []loginErrorCase{
		loginErrorCase{
			email:    "t@t.com",
			password: "a",
		},
		loginErrorCase{
			email:    "t@t.com",
			password: "",
		},
		loginErrorCase{
			email:    "",
			password: "a",
		},
		loginErrorCase{
			email:    "",
			password: "",
		},
		loginErrorCase{
			email:    "t@t.com;DROP users;--", // Check for the most obvious SQLi attempt.
			password: "",
		},
		loginErrorCase{
			email:    "t",
			password: "abc123",
		},
	}

	for _, tC := range testCases {
		if err := api.Login(tC.email, tC.password); err != client.ErrIncorrectEmailOrPassword {
			t.Errorf("'%s', '%s': Expected %v, Got %v", tC.email, tC.password,
				client.ErrIncorrectEmailOrPassword, err)
		}
	}

	if err := api.Login(email, password); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user, err := api.Whoami()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if user.Email != email {
		t.Errorf("Expected %s, Got %s", email, user.Email)
	}
}
