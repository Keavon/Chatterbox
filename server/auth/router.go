package auth

import "github.com/gorilla/mux"

// New initalizes a new auth router.
func New(r *mux.Router) {
	r.Path("/auth").Methods("POST").HandlerFunc(register)
	r.Path("/auth/login").Methods("POST").HandlerFunc(login)
	r.Path("/auth/user").Methods("GET").HandlerFunc(CheckAuth(getUser))
	r.Path("/auth/user").Methods("POST").HandlerFunc(CheckAuth(updateUser))
}
