package auth

import (
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

var invalidToken = util.ErrorRes{Errors: []util.ErrorMsg{util.ErrorMsg{Msg: "invalid token"}}}

// UserReq is a request by an authenticated user.
type UserReq func(http.ResponseWriter, *http.Request, *models.User)

// CheckAuth checks authentication and
func CheckAuth(fn UserReq) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valid, usrID := userToken.Valid(r.Header.Get("Authorization"))

		if !valid {
			util.JSONResponse(w, invalidToken, 401)
			return
		}

		usr, err := models.GetUser(usrID, "")

		if err != nil {
			logger.Error.Print(err)
			w.WriteHeader(500)
			return
		}

		fn(w, r, usr)
	}
}
