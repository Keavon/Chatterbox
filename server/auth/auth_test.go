package auth

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"

	"github.com/chatterbox-irc/chatterbox/server/client"
	"github.com/chatterbox-irc/chatterbox/server/mock"
)

func NewAuthTest(t *testing.T) (*client.Client, *httptest.Server) {
	if err := mock.NewMockDB(); err != nil {
		t.Fatal(err)
	}

	mock.NewMockLogger()

	r := mux.NewRouter()
	New(r)
	server := httptest.NewServer(r)

	u, err := url.Parse(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	return &client.Client{Host: *u}, server
}
