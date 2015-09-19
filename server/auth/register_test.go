package auth

import (
	"reflect"
	"testing"

	"github.com/chatterbox-irc/chatterbox/server/client"
)

type registerErrorCase struct {
	errMsgs  []client.ErrorMsg
	email    string
	password string
}

func TestRegistration(t *testing.T) {
	api, server := NewAuthTest(t)
	defer server.Close()

	if _, err := api.Register("t@t.com", "abc123"); err != nil {
		t.Error(err)
	}

	testCases := []registerErrorCase{
		registerErrorCase{
			errMsgs: []client.ErrorMsg{
				client.ErrorMsg{
					Field: "email",
					Msg:   "not a valid email",
				},
			},
			email:    "h@",
			password: "abc123",
		},
		registerErrorCase{
			errMsgs: []client.ErrorMsg{
				client.ErrorMsg{
					Field: "email",
					Msg:   "can't be nil",
				},
				client.ErrorMsg{
					Field: "email",
					Msg:   "not a valid email",
				},
			},
			email:    "",
			password: "abc123",
		},
		registerErrorCase{
			errMsgs: []client.ErrorMsg{
				client.ErrorMsg{
					Field: "password",
					Msg:   "can't be nil",
				},
			},
			email:    "t@test.com",
			password: "",
		},
		registerErrorCase{
			errMsgs: []client.ErrorMsg{
				client.ErrorMsg{
					Field: "email",
					Msg:   "must be unique",
				},
			},
			email:    "t@t.com",
			password: "abc123",
		},
	}

	for _, tC := range testCases {
		errMsgs, _ := api.Register(tC.email, tC.password)

		if !reflect.DeepEqual(errMsgs, tC.errMsgs) {
			t.Errorf("Expected %v, Got %v", tC.errMsgs, errMsgs)
		}
	}
}
