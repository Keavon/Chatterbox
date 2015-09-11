package models

import (
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// User is the structure of the user model.
type User struct {
	ID       string `sql:"not null;unique" json:"id"`
	Email    string `sql:"not null;unique" json:"email"`
	Password string `sql:"not null" json:"password"`
}

// NewUser creates a new user object if it is valid.
func NewUser(email, password string) (User, []ValidationMsg, error) {
	id, err := generateUUID()
	if err != nil {
		logger.Error.Print(err)
		return User{}, []ValidationMsg{}, err
	}

	msg := ValidateUser(id, email, password)

	if len(msg) > 0 {
		return User{}, msg, nil
	}

	// Use bcrypt to hash password. Using cost 10 (reccomended by library).
	ePass, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		logger.Error.Print(err)
		return User{}, []ValidationMsg{}, err
	}

	user := User{ID: id, Email: email, Password: string(ePass)}

	if err = DB.Create(&user).Error; err != nil {
		logger.Error.Print(err)
		return User{}, []ValidationMsg{}, err
	}

	return user, []ValidationMsg{}, nil
}

// ValidateUser validates user models
func ValidateUser(id, email, password string) (e []ValidationMsg) {
	e = append(e, notNil("id", id)...)
	e = append(e, notNil("email", email)...)
	e = append(e, notNil("password", password)...)

	e = append(e, isEmail("email", email)...)

	if !DB.Where(&User{ID: id}).First(&User{}).RecordNotFound() {
		e = append(e, ValidationMsg{Field: "id", Msg: uniqueMsg})
	}

	if !DB.Where(&User{Email: email}).First(&User{}).RecordNotFound() {
		e = append(e, ValidationMsg{Field: "email", Msg: uniqueMsg})
	}

	return e
}
