package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

// ErrorRes is a general json error response.
type ErrorRes struct {
	Errors []string `json:"errors"`
}

// ReadBody reads a request body and handles errors.
func ReadBody(b io.Reader, w http.ResponseWriter) (string, error) {
	body, err := ioutil.ReadAll(b)

	if err != nil {
		logger.Error.Print(err)
		w.WriteHeader(500)
		return "", err
	}

	return string(body), nil
}

// ParseJSON parses json in body and marshals it to a struct
func ParseJSON(b io.Reader, w http.ResponseWriter, m interface{}) error {
	body, err := ReadBody(b, w)

	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(body), m); err != nil {
		logger.Debug.Print(err)
		JSONResponse(w, ErrorRes{Errors: []string{"invalid json"}}, 400)
		return err
	}

	return nil
}

// JSONResponse writes a json response from a struct
func JSONResponse(w http.ResponseWriter, b interface{}, responseCode int) {
	res, err := json.Marshal(b)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	w.Write(res)
}
