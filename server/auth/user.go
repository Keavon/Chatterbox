package auth

import (
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

func getUser(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.WriteHeader(200)
	util.JSONResponse(w, u, 200)
}
