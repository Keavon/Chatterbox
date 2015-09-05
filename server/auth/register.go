package auth

import (
	"fmt"
	"net/http"

	"github.com/chatterbox-irc/chatterbox/server/models"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
	"github.com/chatterbox-irc/chatterbox/server/util"
)

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var err error
	logger.Debug.Printf("%s %s\n", r.Method, r.URL.Path)

	user := models.User{}

	if err := util.ParseJSON(r.Body, w, &user); err != nil {
		return
	}

	if user.ID, err = util.GenerateUUID(w); err != nil {
		return
	}

	if err != nil {
		return
	}

	logger.Debug.Println(user.Validate())

	fmt.Fprint(w, r.URL.Path[1:])
}
