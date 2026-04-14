package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jarusuraj/schoolsystem/config"
)

func InsertSubjects(subjects []string) error {
	if len(subjects) == 0 {
		return nil
	}

	query := "INSERT INTO subjects (name) VALUES "
	values := []string{}
	args := []interface{}{}

	for i, s := range subjects {
		values = append(values, fmt.Sprintf("($%d)", i+1))
		args = append(args, s)
	}

	query += strings.Join(values, ", ")
	query += " ON CONFLICT (name) DO NOTHING"

	_, err := config.DB.Exec(context.Background(), query, args...)
	return err
}
