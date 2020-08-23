package oauth2

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var errWrongPassword = errors.New("Password is wrong")

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (o *OAuth2) verifyUser(mail string, password string) (string, error) {
	user, err := o.users.Repository.GetByMail(mail)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(password, user.Password) {
		return "", errWrongPassword
	}

	return user.ID.String(), nil
}
