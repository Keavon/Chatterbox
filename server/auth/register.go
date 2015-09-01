package auth

import (
	"fmt"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path[1:])
}
