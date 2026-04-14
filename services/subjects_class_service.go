package services

import (
	"github.com/jarusuraj/schoolsystem/models"
	"github.com/jarusuraj/schoolsystem/repository"
)

func GetSubjects() ([]models.SubjectRequest, error) {
	return repository.GetAllSubjects()
}
func AddClassSubjects(req models.BulkClassSubjectsRequest, classNo int) error {
	return repository.InsertClassSubjects(req.Items, classNo)
}