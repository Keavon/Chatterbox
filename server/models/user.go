package models

import (
	"github.com/chatterbox-irc/chatterbox/pkg/validate"
	"github.com/chatterbox-irc/chatterbox/server/pkg/logger"
	"github.com/jinzhu/gorm"

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
func (u *User) Update(email, password string) ([]validate.ValidationMsg, error) {
	updated := false

	if password != "" {
		ePass, err := bcryptPass(password)

		if err != nil {
			return []validate.ValidationMsg{}, err
		}

		updated = true
		u.Password = ePass
	}

	if email != "" {
		e := checkEmail(email, u.ID)

		if len(e) > 0 {
			return e, nil
		}

		updated = true
		u.Email = email
	}

	if updated {
		if err := DB.Save(u).Error; err != nil {
			return []validate.ValidationMsg{}, err
		}
	}

	return []validate.ValidationMsg{}, nil
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
func NewUser(email, password string) (*User, []validate.ValidationMsg, error) {
	id, err := generateUUID()
	if err != nil {
		logger.Error.Print(err)
		return &User{}, []validate.ValidationMsg{}, err
	}

	msg := ValidateUser(id, email, password)

	if len(msg) > 0 {
		return &User{}, msg, nil
	}

	ePass, err := bcryptPass(password)

	if err != nil {
		return &User{}, []validate.ValidationMsg{}, err
	}

	user := User{ID: id, Email: email, Password: ePass}

	if err = DB.Create(&user).Error; err != nil {
		logger.Error.Print(err)
		return &User{}, []validate.ValidationMsg{}, err
	}

	return &user, []validate.ValidationMsg{}, nil
}

// ValidateUser validates user models
func ValidateUser(id, email, password string) (e []validate.ValidationMsg) {
	e = append(e, validate.NotNil("id", id)...)
	e = append(e, validate.NotNil("email", email)...)
	e = append(e, validate.NotNil("password", password)...)

	e = append(e, checkEmail(email, "")...)

	if !DB.Where(&User{ID: id}).First(&User{}).RecordNotFound() {
		e = append(e, validate.ValidationMsg{Field: "id", Msg: validate.UniqueMsg})
	}

	return e
}

func checkEmail(email string, id string) (e []validate.ValidationMsg) {
	e = append(e, validate.IsEmail("email", email)...)

	if len(e) > 0 {
		return
	}

	user := User{}

	if !DB.Where(&User{Email: email}).First(&user).RecordNotFound() && user.ID != id {
		e = append(e, validate.ValidationMsg{Field: "email", Msg: validate.UniqueMsg})
	}
	return
}

// GetUser retrives a user by email
func GetUser(id, email string) (*User, error) {
	user := User{}

	if id != "" {
		return &user, DB.Where("id = ?", id).First(&user).Error
	} else if email != "" {
		return &user, DB.Where("email = ?", email).First(&user).Error
	}

	return &User{}, gorm.RecordNotFound
}
