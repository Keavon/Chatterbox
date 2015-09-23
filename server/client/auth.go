package client

import (
	"encoding/json"
	"fmt"
)

// User is the definition of the user object
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type loginRes struct {
	Token string `json:"token"`
}

// Register a new user.
func (c Client) Register(email, password string) ([]ErrorMsg, error) {
	user := User{Email: email, Password: password}

	req, err := json.Marshal(user)

	if err != nil {
		return []ErrorMsg{}, err
	}

	body, _, err := c.Request("POST", "/auth", req)

	if err != ErrValidation {
		return []ErrorMsg{}, err
	}

	res := ErrorRes{}

	if parseErr := json.Unmarshal(body, &res); parseErr != nil {
		return []ErrorMsg{}, err
	}

	return res.Errors, err
}

// Login and retrieve token.
func (c *Client) Login(email, password string) error {
	user := User{Email: email, Password: password}

	req, err := json.Marshal(user)

	if err != nil {
		return err
	}

	body, _, err := c.Request("POST", "/auth/login", req)

	if err != nil {
		return err
	}

	res := loginRes{}

	if err = json.Unmarshal(body, &res); err != nil {
		return err
	}

	c.Token = res.Token

	return nil
}

// Whoami returns authorized user.
func (c Client) Whoami() (*User, error) {
	body, _, err := c.Request("GET", "/auth/user", []byte{})

	if err != nil {
		return &User{}, err
	}

	res := User{}

	if err = json.Unmarshal(body, &res); err != nil {
		return &User{}, err
	}

	return &res, nil
}

// UpdateUser updates an user.
func (c *Client) UpdateUser(user *User) (*User, error) {

	req, err := json.Marshal(user)

	if err != nil {
		return &User{}, err
	}

	fmt.Println("-------> ", string(req))

	body, _, err := c.Request("POST", "/auth/user", req)

	if err != nil {
		fmt.Println(">>>>>>>>>> Err: ", err.Error())
		return &User{}, err
	}

	res := User{}

	if err = json.Unmarshal(body, &res); err != nil {
		return &User{}, err
	}

	return &res, nil
}
