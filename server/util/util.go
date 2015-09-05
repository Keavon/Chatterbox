package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/wayn3h0/go-uuid"
	"github.com/wayn3h0/go-uuid/random"

	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

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
		logger.Error.Print(err)
		w.WriteHeader(400)
		w.Write([]byte(`{"errors": ["invalid json"]}`))
		return err
	}

	return nil
}

// GenerateUUID generates a uuid
func GenerateUUID(w http.ResponseWriter) (string, error) {
	nUUID, err := random.New()

	if err != nil {
		logger.Error.Print(err)
		w.WriteHeader(500)
		w.Write([]byte(`{"errors": ["unable to generate random number"]}`))
		return "", err
	}

	return nUUID.Format(uuid.StyleWithoutDash), nil
}
