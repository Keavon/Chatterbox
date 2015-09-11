package auth

import (
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	logger.Debug.Printf("%s %s\n", r.Method, r.URL.Path)

	req := registerReq{}

	if err := util.ParseJSON(r.Body, w, &req); err != nil {
		// ParseJSON handles error reponse
		return
	}

	_, msg, err := models.NewUser(req.Email, req.Password)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if len(msg) > 0 {
		msgJSON, err := models.ValidationToJSON(msg)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(400)
		w.Write(msgJSON)
	}

	w.WriteHeader(201)
}
