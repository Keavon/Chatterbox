package irc

import (
	"fmt"

	ircevent "github.com/thoj/go-ircevent"

	"github.com/chatterbox-irc/chatterbox/ircc/events"
)

// OnConnect runs after connecting to an IRC server.
func (i *IRC) OnConnect(e *ircevent.Event) {
	fmt.Fprintln(i.Output, events.Connected(i.Server))
	i.Connected = true
}

// OnJoin runs after joining a channel.
func (i *IRC) OnJoin(e *ircevent.Event) {
	fmt.Fprintln(i.Output, events.Joined(e.Arguments[0]))
}

// OnPart runs after parting a channel.
func (i *IRC) OnPart(e *ircevent.Event) {
	fmt.Fprintln(i.Output, events.Parted(e.Arguments[0]))
}

// OnMsg runs on message arrival.
func (i IRC) OnMsg(e *ircevent.Event) {
	notice := false
	if e.Code == "NOTICE" {
		notice = true
	}

	fmt.Fprintln(i.Output, events.RcvedMsg(e.Nick, e.Arguments[0], e.Arguments[1], notice))
}
