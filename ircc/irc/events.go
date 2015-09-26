package irc

import ircevent "github.com/thoj/go-ircevent"

// OnConnect runs after connecting to an IRC server.
func (i *IRC) OnConnect(e *ircevent.Event) {
	i.Connected = true
}
