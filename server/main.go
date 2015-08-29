package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chatterbox-irc/chatterbox/server/auth"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = 8080
)

var (
	host   = flag.String("host", defaultHost, "Host to bind to.")
	port   = flag.Int("port", defaultPort, "Port to bind to.")
	listen = fmt.Sprintf("%s:%d", *host, *port)
)

func main() {
	fmt.Printf("Chatterbox listening on %s.\n", listen)
	r := mux.NewRouter()
	auth.New(r)

	http.Handle("/", r)
	http.ListenAndServe(listen, nil)
}
