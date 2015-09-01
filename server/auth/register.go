package auth

import (
	"fmt"
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

func register(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Printf("%s %s\n", r.Method, r.URL.Path)

	fmt.Fprint(w, r.URL.Path[1:])
}
