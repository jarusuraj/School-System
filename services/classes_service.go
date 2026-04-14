package services

import (
	"errors"

	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/repository"
)

func CreateClasses(req models.BulkCreateClassRequest) error {
	if len(req.Classes) == 0 {
		return errors.New("no classes provided")
	}

	return repository.InsertClasses(req.Classes)
}
