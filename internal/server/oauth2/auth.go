package oauth2

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/archstack/workspace-api/lib/models"
)

var errWrongPassword = errors.New("Password is wrong")

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (o *OAuth2) verifyUser(user *models.User, password string) error {
	if !checkPasswordHash(password, user.Password) {
		return errWrongPassword
	}

	return nil
}
