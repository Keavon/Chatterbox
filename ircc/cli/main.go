package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chatterbox-irc/chatterbox/ircc/events"
	"github.com/chatterbox-irc/chatterbox/ircc/irc"
	"github.com/chatterbox-irc/chatterbox/ircc/parser"
)

const (
	connectionTimeout = time.Second * 10
)

var (
	nick   = flag.String("nick", "", "Nickname")
	user   = flag.String("user", "", "User full name")
	server = flag.String("server", "", "Server address")
	pass   = flag.String("pass", "", "Server password")
	tls    = flag.Bool("tls", false, "Connect using tls")
)

type connectionError struct {
}

func main() {
	flag.Parse()
	connect(os.Stdout, os.Stdin)
}

func connect(w io.Writer, r io.Reader) {
	ircc, err := irc.New(*nick, *user, *server, *pass, *tls, w)
	if err != nil {
		fmt.Fprintln(w, events.ConnectionError(err.Error()))
		os.Exit(1)
	}

	err = ircc.WaitForConnection(connectionTimeout)

	if err != nil {
		os.Exit(2)
	}

	inReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inReader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(w, events.InternalError(err.Error()))
		}

		status := parser.Parse(ircc, w, input)

		if status != -1 {
			os.Exit(status)
		}
	}
}
