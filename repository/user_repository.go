package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	config "github.com/jarusuraj/schoolsystem/config"
)

var db = config.DB

func Signup(email_id string, hash []byte, name, role, phone string) (err error) {

	chcek := "SELECT (email) from users where email=($1)"
	var email string
	err = db.QueryRow(context.Background(), chcek, email_id).Scan(&email)
	switch err {
	case nil:
		return errors.New("Email Already Exists")
	case pgx.ErrNoRows:
		query := "INSERT INTO users (name,email,password,role,phone) VALUES ($1,$2,$3,$4,$5)"
		_, _ = db.Exec(context.Background(), query, email_id, hash, name, role, phone)
		return nil
	default:
		return err
	}
}
func Login(email_id string) ([]byte, error) {
	var hashPassword []byte
	if err := db.QueryRow(context.Background(),
		"SELECT password_hash FROM users WHERE email=$1",
		email_id,
	).Scan(&hashPassword); err != nil {
		return nil, err
	}
	return hashPassword, nil
}
