package auth

import (
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := registerReq{}

	if err := util.ParseJSON(r.Body, w, &req); err != nil {
		return
	}

	_, msg, err := models.NewUser(req.Email, req.Password)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if len(msg) > 0 {
		util.JSONResponse(w, models.ValidationToJSON(msg), 400)
	}

	w.WriteHeader(201)
}
