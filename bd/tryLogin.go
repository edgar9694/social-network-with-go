package bd

import (
	"github.com/edgar9694/social-network-with-go.git/models"
	"golang.org/x/crypto/bcrypt"
)

/* TryLogin() realiza el chequeo de login a la BD*/
func TryLogin(email string, password string) (models.User, bool) {
	usu, finded, _ := CheckUserExists(email)
	if !finded {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true

}
