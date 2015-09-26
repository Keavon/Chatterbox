package auth

import (
	"net/http"

	"github.com/chatterbox-irc/chatterbox/pkg/validate"
	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

func getUser(w http.ResponseWriter, r *http.Request, u *models.User) {
	util.JSONResponse(w, u, 200)
}

func updateUser(w http.ResponseWriter, r *http.Request, u *models.User) {
	defer r.Body.Close()

	req := registerReq{}

	if err := util.ParseJSON(r.Body, w, &req); err != nil {
		return
	}

	msg, err := u.Update(req.Email, req.Password)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if len(msg) > 0 {
		util.JSONResponse(w, validate.ValidationMsgsToJSON(msg), 400)
		return
	}

	util.JSONResponse(w, u, 200)
}
