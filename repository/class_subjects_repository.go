package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jarusuraj/schoolsystem/config"
	"github.com/jarusuraj/schoolsystem/models"
)

func GetAllSubjects() ([]models.SubjectRequest, error) {
	query := "SELECT id, name FROM subjects"

	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []models.SubjectRequest

	for rows.Next() {
		var s models.SubjectRequest

		err := rows.Scan(&s.ID, &s.SubjectName)
		if err != nil {
			return nil, err
		}

		subjects = append(subjects, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}
func InsertClassSubjects(items []models.ClassSubjectRequest, classID int) error {
	values := []string{}
	args := []interface{}{}

	for i, item := range items {
		base := i*3 + 1

		values = append(values,
			fmt.Sprintf("($%d,$%d,$%d)", base, base+1, base+2),
		)

		args = append(args,
			classID,
			item.SubjectID,
			item.SubjectCode,
		)
	}

	query := `
		INSERT INTO class_subjects (class_no, subject_id, subject_code)
		VALUES ` + strings.Join(values, ", ")

	_, err := config.DB.Exec(context.Background(), query, args...)
	return err
}