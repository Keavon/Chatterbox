package auth

import (
	"net/http"
	"time"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRes struct {
	Token string `json:"token"`
}

var incorectEmailOrPassword = util.ErrorRes{Errors: []util.ErrorMsg{util.ErrorMsg{Msg: "incorrect email or password"}}}

func login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := loginReq{}

	if err := util.ParseJSON(r.Body, w, &req); err != nil {
		// ParseJSON handles error reponse
		return
	}

	user, err := models.GetUser("", req.Email)

	if err != nil {
		logger.Debug.Println("User not found")
		util.JSONResponse(w, incorectEmailOrPassword, 401)
		return
	}

	if !user.CheckPass(req.Password) {
		logger.Debug.Println("Incorrect Password")
		util.JSONResponse(w, incorectEmailOrPassword, 401)
		return
	}

	token, err := userToken.New(user.ID, time.Now().Add(userTokenExp))

	if err != nil {
		logger.Error.Print(err)
		w.WriteHeader(500)
		return
	}

	util.JSONResponse(w, loginRes{Token: token}, 200)
}
