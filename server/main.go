package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/chatterbox-irc/chatterbox/server/auth"
	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
)

const (
	defaultHost      = "0.0.0.0"
	defaultPort      = 8080
	defaultLog       = "console"
	defaultLogLevel  = "debug"
	dbConnectionHelp = "See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters"
)

var (
	host      = flag.String("host", defaultHost, "Host to bind to.")
	port      = flag.Int("port", defaultPort, "Port to bind to.")
	db        = flag.String("db", "", "DB connection string\n"+dbConnectionHelp)
	logOutput = flag.String("log", defaultLog, "Log streaming location. Valid options: console")
	logLevel  = flag.String("loglevel", defaultLogLevel,
		"Log verbosity. Valid options: debug, info, warn, error")
	migrate = flag.Bool("migrate", false, "migrate the database and exit")
)

func main() {
	flag.Parse()
	listen := fmt.Sprintf("%s:%d", *host, *port)

	lo, ll, err := logger.StringToFlag(*logOutput, *logLevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.New(lo, ll)

	if *db == "" {
		logger.Error.Println("A db connection string (flag --db='') is required.\n" + dbConnectionHelp)
		os.Exit(1)
	}

	if err = models.New(*db); err != nil {
		logger.Error.Println(err)
		os.Exit(1)
	}

	if *migrate {
		logger.Info.Println("Migrating Database.")
		models.Migrate()
		os.Exit(0)
	}

	r := mux.NewRouter()
	auth.New(r)

	http.Handle("/", r)
	logger.Info.Printf("Chatterbox listening on %s.\n", listen)
	http.ListenAndServe(listen, nil)
}
