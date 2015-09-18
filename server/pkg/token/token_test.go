package token

import (
	"testing"
	"time"
)

var (
	testToken      = Token{Secret: "abc123", ISS: "test"}
	testOtherToken = Token{Secret: "123abc", ISS: "other-test"}
)

func TestEncodeDecode(t *testing.T) {
	expected := "test"
	token, err := testToken.New(expected, time.Now().Add(time.Hour))

	if err != nil {
		t.Fatal(err)
	}

	valid, id := testToken.Valid(token)

	if !valid {
		t.Error("Expected token to be valid.")
	}

	if id != expected {
		t.Errorf("Expected %s, Got %s", expected, id)
	}
}

func TestEncodeDecodeFail(t *testing.T) {
	expected := "test"
	token, err := testOtherToken.New(expected, time.Now().Add(time.Hour))

	if err != nil {
		t.Fatal(err)
	}

	valid, _ := testToken.Valid(token)

	if valid {
		t.Error("Expected token to be invalid.")
	}
}
