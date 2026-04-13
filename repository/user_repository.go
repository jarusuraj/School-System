package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	config "github.com/jarusuraj/schoolsystem/config"
	"github.com/jarusuraj/schoolsystem/models"
)

var db = config.DB

func CreateUser(email_id string, hash []byte, name, role, phone string) (err error) {

	chcek := "SELECT (email) from users where email=($1)"
	var email string
	var status string = "pending"
	err = db.QueryRow(context.Background(), chcek, email_id).Scan(&email)
	switch err {
	case nil:
		return errors.New("Email Already Exists")
	case pgx.ErrNoRows:
		query := "INSERT INTO users (name,email,password,role,phone,status) VALUES ($1,$2,$3,$4,$5,$6)"
		_, _ = db.Exec(context.Background(), query, email_id, hash, name, role, phone, status)
		return nil
	default:
		return err
	}
}
func GetUserByEmail(email_id string) (models.User, error) {
	var user models.User
	if err := db.QueryRow(context.Background(),
		"SELECT id,name,email,password,role,status FROM users WHERE email=$1",
		email_id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role,user.Status); err != nil {
		return models.User{}, err
	}
	return user, nil
}
