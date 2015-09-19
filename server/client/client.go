package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is a object for interacting with the API.
type Client struct {
	Host  url.URL
	Token string
}

// ErrorRes is a general json error response.
type ErrorRes struct {
	Errors []ErrorMsg `json:"errors"`
}

// ErrorMsg is a general error msg.
type ErrorMsg struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

var (
	// ErrUnauthorized is a error thrown when authentication fails
	ErrUnauthorized = errors.New("unauthorized")
	// ErrIncorrectEmailOrPassword is a error thrown when login fails.
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")
	// ErrInvalidJSON is thrown when the server recieves invalid json.
	ErrInvalidJSON = errors.New("invalid json")
	// ErrValidation is thrown when server validation of an object fails.
	ErrValidation = errors.New("validation failed")
)

// Request makes a http request.
func (c Client) Request(method string, path string, body []byte) ([]byte, int, error) {
	c.Host.Path = path
	httpClient := &http.Client{}

	req, err := http.NewRequest(method, c.Host.String(), bytes.NewBuffer(body))

	if err != nil {
		return []byte{}, -1, err
	}

	if c.Token != "" {
		req.Header.Add("Authorization", c.Token)
	}

	req.Header.Add("User-Agent", "CBX Go Library")

	res, err := httpClient.Do(req)

	if err != nil {
		return []byte{}, -1, err
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte{}, -1, err
	}

	if res.StatusCode < 200 || res.StatusCode > 399 {
		var err error

		switch res.StatusCode {
		case 401:
			err = ErrUnauthorized
		case 400:
			reqErrs := ErrorRes{}

			if err = json.Unmarshal(body, &reqErrs); err == nil {
				for _, reqErr := range reqErrs.Errors {
					switch reqErr.Msg {
					case "incorrect email or password":
						err = ErrIncorrectEmailOrPassword
						break
					case "invalid json":
						err = ErrInvalidJSON
						break
					}
				}

				// If error didn't match, it's a validation error
				if err == nil {
					err = ErrValidation
				}
			}
		default:
			err = errors.New(string(body))
		}

		return body, res.StatusCode, err
	}

	return body, res.StatusCode, nil
}
