package services

import (
	"errors"
	"strings"

	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/repository"
)

func AddSubjects(subjects []models.SubjectRequest) ([]string, error) {
	if len(subjects) == 0 {
		return nil, errors.New("no subjects provided")
	}

	clean := make([]string, 0, len(subjects))
	seen := make(map[string]bool)

	for _, s := range subjects {
		name := strings.TrimSpace(s.SubjectName)

		if name == "" {
			continue
		}

		key := strings.ToLower(name)

		if seen[key] {
			continue
		}

		seen[key] = true
		clean = append(clean, name)
	}

	if len(clean) == 0 {
		return nil, errors.New("no valid subjects provided")
	}

	err := repository.InsertSubjects(clean)
	if err != nil {
		return nil, err
	}

	return clean, nil
}
