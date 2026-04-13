package services

import (
	"errors"

	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/repository"
	"golang.org/x/crypto/bcrypt"
)

func Signup(user models.SignupRequest) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	if err := repository.CreateUser(user.Email, hashPassword, user.Name, user.Role, user.Phone); err != nil {
		return err
	}
	return nil
}
func Login(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	switch user.Status {
	case "pending":
		return "", errors.New("user not verified. Status Pending!!!")
	case "rejected":
		return "", errors.New("user rejected")
	case "verified":
		err = bcrypt.CompareHashAndPassword(
			[]byte(user.PasswordHash),
			[]byte(password),
		)
		if err != nil {
			return "", err
		}
		token, err := GenerateToken(user.Email, user.ID, user.Role)

		return token, err
	default:
		return "", errors.New("invalid user status")
	}
}
