package models

import (
	"fmt"
)

// User is the structure of the user model.
type User struct {
	ID       string `sql:"not null;unique" json:"id"`
	Email    string `sql:"not null;unique" json:"email"`
	Password string `sql:"not null" json:"password"`
}

// Validate User models
func (u User) Validate() (e []string) {
	e = append(e, notNil("ID", u.ID)...)
	e = append(e, notNil("Email", u.Email)...)
	e = append(e, notNil("Password", u.Password)...)

	e = append(e, isEmail(u.Email)...)

	if !DB.Where(&User{ID: u.ID}).First(&User{}).RecordNotFound() {
		e = append(e, fmt.Sprintf(uniqueErrTemplate, "ID"))
	}

	if !DB.Where(&User{Email: u.Email}).First(&User{}).RecordNotFound() {
		e = append(e, fmt.Sprintf(uniqueErrTemplate, "Email"))
	}

	return e
}
