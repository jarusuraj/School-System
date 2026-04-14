package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jarusuraj/schoolsystem/config"
	"github.com/jarusuraj/schoolsystem/models"
)

func InsertClasses(classes []models.CreateClassRequest) error {
	if len(classes) == 0 {
		return nil
	}

	values := []string{}
	args := []interface{}{}

	for i, c := range classes {
		base := i*2 + 1

		values = append(values,
			fmt.Sprintf("($%d,$%d)", base, base+1),
		)

		args = append(args,
			c.ClassNo,
			c.Name,
		)
	}

	query := `
		INSERT INTO classes (class_no, name)
		VALUES ` + strings.Join(values, ",")

	_, err := config.DB.Exec(context.Background(), query, args...)
	return err
}
