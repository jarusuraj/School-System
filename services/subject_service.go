package services

import (
	"errors"
	"strings"

	"github.com/jarusuraj/schoolsystem/repository"
)

func AddSubjects(subjects []string) error {
	if len(subjects) == 0 {
		return errors.New("No Subjects are provided")
	}
	for i := range subjects {
		subjects[i] = strings.TrimSpace(subjects[i])
	}
	return repository.InsertSubjects(subjects)
}
