package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// New initalizes a new auth router.
func New(r *mux.Router) {
	a := r.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/", test)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path[1:])
}
