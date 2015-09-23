package auth

import "testing"

func TestUser(t *testing.T) {
	api, server := NewAuthTest(t)
	defer server.Close()

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

	user.Email = "test@test.com"
	user, err = api.UpdateUser(user)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Can't log in with old email
	if err := api.Login(email, password); err == nil {
		t.Error("Expected error")
	}

	email = "test@test.com"

	if err := api.Login(email, password); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	user.Password = "123abc"
	user, err = api.UpdateUser(user)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Can't log in with old password
	if err := api.Login(email, password); err == nil {
		t.Error("Expected error")
	}

	password = "123abc"

	if err := api.Login(email, password); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
