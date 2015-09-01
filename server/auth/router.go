package auth

import "github.com/gorilla/mux"

// New initalizes a new auth router.
func New(r *mux.Router) {
	a := r.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/", register).Methods("POST")
}
