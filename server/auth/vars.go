package auth

import (
	"os"
	"time"

	"github.com/chatterbox-irc/chatterbox/server/pkg/token"
)

var userToken = token.Token{ISS: "cbx.user", Secret: os.Getenv("USER_TOKEN_SECRET")}
var userTokenExp = time.Duration(72 * time.Hour)
