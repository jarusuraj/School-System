package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	config "github.com/jarusuraj/schoolsystem/config"
	"github.com/jarusuraj/schoolsystem/models"
)



func CreateUser(email string, hash []byte, name, role, phone string) error {

	var exists string
	check := "SELECT email FROM users WHERE email = $1"

	err := config.DB.QueryRow(context.Background(), check, email).Scan(&exists)
	if err == nil {
		return errors.New("email already exists")
	}

	if err != pgx.ErrNoRows {
		return err
	}

	status := "pending"

	query := `
		INSERT INTO users (name, email, password, role, phone, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = config.DB.Exec(context.Background(),
		query,
		name,
		email,
		hash,
		role,
		phone,
		status,
	)

	return err
}
func GetUserByEmail(email_id string) (models.User, error) {
	var user models.User
	if err := config.DB.QueryRow(context.Background(),
		"SELECT id,name,email,password,role,status FROM users WHERE email=$1",
		email_id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.Status); err != nil {
		return models.User{}, err
	}
	return user, nil
}
