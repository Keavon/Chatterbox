package models

import (
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// User is the structure of the user model.
type User struct {
	ID       string `sql:"not null;unique" json:"id"`
	Email    string `sql:"not null;unique" json:"email"`
	Password string `sql:"not null" json:"-"`
}

// CheckPass checks if the password matches the users hash.
func (u User) CheckPass(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err == nil {
		return true
	}

	if err != bcrypt.ErrMismatchedHashAndPassword {
		logger.Error.Print(err)
	}

	return false
}

// Update updates an existing user.
func (u *User) Update(email, password string) ([]ValidationMsg, error) {
	updated := false

	if password != "" {
		ePass, err := bcryptPass(password)

		if err != nil {
			return []ValidationMsg{}, err
		}

		updated = true
		u.Password = ePass
	}

	if email != "" {
		e := checkEmail(email)

		if len(e) > 0 {
			return e, nil
		}

		updated = true
		u.Email = email
	}

	if updated {
		if err := DB.Save(u).Error; err != nil {
			return []ValidationMsg{}, err
		}
	}

	return []ValidationMsg{}, nil
}

func bcryptPass(password string) (string, error) {
	// Use bcrypt to hash password. Using cost 10 (reccomended by library).
	ePass, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		logger.Error.Print(err)
		return "", err
	}

	return string(ePass), nil
}

// NewUser creates a new user object if it is valid.
func NewUser(email, password string) (*User, []ValidationMsg, error) {
	id, err := generateUUID()
	if err != nil {
		logger.Error.Print(err)
		return &User{}, []ValidationMsg{}, err
	}

	msg := ValidateUser(id, email, password)

	if len(msg) > 0 {
		return &User{}, msg, nil
	}

	ePass, err := bcryptPass(password)

	if err != nil {
		return &User{}, []ValidationMsg{}, err
	}

	user := User{ID: id, Email: email, Password: ePass}

	if err = DB.Create(&user).Error; err != nil {
		logger.Error.Print(err)
		return &User{}, []ValidationMsg{}, err
	}

	return &user, []ValidationMsg{}, nil
}

// ValidateUser validates user models
func ValidateUser(id, email, password string) (e []ValidationMsg) {
	e = append(e, notNil("id", id)...)
	e = append(e, notNil("email", email)...)
	e = append(e, notNil("password", password)...)

	e = append(e, checkEmail(email)...)

	if !DB.Where(&User{ID: id}).First(&User{}).RecordNotFound() {
		e = append(e, ValidationMsg{Field: "id", Msg: uniqueMsg})
	}

	return e
}

func checkEmail(email string) (e []ValidationMsg) {
	e = append(e, isEmail("email", email)...)

	if !DB.Where(&User{Email: email}).First(&User{}).RecordNotFound() {
		e = append(e, ValidationMsg{Field: "email", Msg: uniqueMsg})
	}
	return
}

// GetUser retrives a user by email
func GetUser(id, email string) (*User, error) {
	user := User{ID: id, Email: email}
	return &user, DB.First(&user).Error
}
