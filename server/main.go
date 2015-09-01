package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/chatterbox-irc/chatterbox/server/auth"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

const (
	defaultHost     = "0.0.0.0"
	defaultPort     = 8080
	defaultLog      = "console"
	defaultLogLevel = "debug"
)

var (
	host      = flag.String("host", defaultHost, "Host to bind to.")
	port      = flag.Int("port", defaultPort, "Port to bind to.")
	logOutput = flag.String("log", defaultLog, "Log streaming location. Valid options: console")
	logLevel  = flag.String("loglevel", defaultLogLevel,
		"Log verbosity. Valid options: debug, info, warn, error")
	listen = fmt.Sprintf("%s:%d", *host, *port)
)

func main() {
	lo, ll, err := logger.StringToFlag(*logOutput, *logLevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.New(lo, ll)

	logger.Info.Printf("Chatterbox listening on %s.\n", listen)
	r := mux.NewRouter()
	auth.New(r)

	http.Handle("/", r)
	http.ListenAndServe(listen, nil)
}
