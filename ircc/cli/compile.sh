#!/usr/bin/env bash
export GO15VENDOREXPERIMENT=1
export DEBIAN_FRONTEND=noninteractive
CGO_ENABLED=0 go build -o $GOPATH/src/github.com/chatterbox-irc/chatterbox/ircc/cli/ircc github.com/chatterbox-irc/chatterbox/ircc/cli
