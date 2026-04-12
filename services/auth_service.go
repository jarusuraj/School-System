package services

import (
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/repository"
	"golang.org/x/crypto/bcrypt"
)

func Signup(user models.SignupRequest) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	if err := repository.Signup(user.Email, hashPassword, user.Name, user.Role, user.Phone); err != nil {
		return err
	}
	return nil
}
func Login(email, password string) (string, error) {
	hashPassword, err := repository.Login(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(hashPassword),
		[]byte(password),
	)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(email)

	return token, err
}
