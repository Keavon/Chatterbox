package irc

// Join an IRC channel.
func (i IRC) Join(channel, password string) {
	con := channel

	if password != "" {
		con += " " + password
	}

	i.Connection.Join(con)
}

// Part from a channel.
func (i *IRC) Part(channel string) {
	i.Connection.Part(channel)
}

// Disconnect from an IRC server.
func (i *IRC) Disconnect() {
	i.Connected = false
	// TODO: Does this fully clean up all go routines?
	// Disconnect() hangs and isn't used in tests.
	i.Connection.Quit()
}

// Msg sends a message to a user or channel
func (i *IRC) Msg(channel, msg string, notice bool) {
	if notice {
		i.Connection.Notice(channel, msg)
		return
	}

	i.Connection.Privmsg(channel, msg)
}
