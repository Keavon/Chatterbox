package irc

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	ircevent "github.com/thoj/go-ircevent"

	"github.com/chatterbox-irc/chatterbox/ircc/events"
)

// IRC is a object holding irc connection information
type IRC struct {
	Nick       string
	User       string
	UseTLS     bool
	Server     string
	ServerPass string
	Output     io.Writer
	Connection *ircevent.Connection
	Connected  bool
}

// New creates a new irc connection
func New(nick, user, server, serverPass string, useTLS bool, output io.Writer) (*IRC, error) {
	con := ircevent.IRC(nick, user)
	con.Log = log.New(ioutil.Discard, "", log.LstdFlags)
	con.UseTLS = useTLS
	con.Password = serverPass

	if err := con.Connect(server); err != nil {
		return nil, err
	}

	go con.Loop()

	ircc := IRC{
		Nick:       nick,
		User:       user,
		UseTLS:     useTLS,
		Server:     server,
		ServerPass: serverPass,
		Output:     output,
		Connection: con,
		Connected:  false,
	}

	con.AddCallback("001", ircc.OnConnect)
	con.AddCallback("JOIN", ircc.OnJoin)
	con.AddCallback("PART", ircc.OnPart)

	return &ircc, nil
}

// WaitForConnection blocks until the irc connection succeeds or times out
func (i *IRC) WaitForConnection(timeout time.Duration) error {
	start := time.Now()

	for {
		if i.Connected == true {
			return nil
		}

		if time.Since(start) > timeout {
			err := errors.New("connection timeout")
			fmt.Fprintln(i.Output, events.ConnectionError(err.Error()))
			return err
		}

		fmt.Fprintln(i.Output, events.WaitForConnection(time.Since(start).Seconds()))
		time.Sleep(500 * time.Millisecond)
	}
}
