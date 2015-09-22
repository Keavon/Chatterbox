package auth

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/chatterbox-irc/chatterbox/server/mock"
	"github.com/chatterbox-irc/chatterbox/server/models"
)

func TestHandler(t *testing.T) {
	api, server := NewAuthTest(t)
	defer server.Close()

	mockResponse := mock.ResponseWriter{Status: 200}
	mockRequest, err := http.NewRequest("GET", "/test", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	testFunc := CheckAuth(func(http.ResponseWriter, *http.Request, *models.User) {
		t.Error("Handler should have failed")
	})

	expected := 401

	testFunc(&mockResponse, mockRequest)
	if mockResponse.Status != expected {
		t.Errorf("Expected %d, Got %s", expected, mockResponse.Status)
	}
	mockResponse.Status = 200

	mockRequest.Header.Set("Authorization", "thiswillfail")
	testFunc(&mockResponse, mockRequest)
	if mockResponse.Status != expected {
		t.Errorf("Expected %d, Got %s", expected, mockResponse.Status)
	}
	mockResponse.Status = 200

	mockRequest.Header.Set("Authorization", "")
	testFunc(&mockResponse, mockRequest)
	if mockResponse.Status != expected {
		t.Errorf("Expected %d, Got %s", expected, mockResponse.Status)
	}
	mockResponse.Status = 200

	email, password := "t@t.com", "abc123"
	if _, err := api.Register(email, password); err != nil {
		t.Error(err)
	}

	if err := api.Login(email, password); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user, err := api.Whoami()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	token, err := userToken.New(user.ID, time.Now().Add(time.Hour))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	testFunc = CheckAuth(func(w http.ResponseWriter, r *http.Request, u *models.User) {
		if u.ID != user.ID {
			t.Errorf("Users not identical: %v %v", u, user)
		}
	})

	expected = 200
	mockRequest.Header.Set("Authorization", token)
	testFunc(&mockResponse, mockRequest)
	if mockResponse.Status != expected {
		t.Errorf("Expected %d, Got %d", expected, mockResponse.Status)
	}
}
