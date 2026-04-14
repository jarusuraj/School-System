package models

type ClassSubjectRequest struct {
	SubjectID   int    `json:"subject_id" binding:"required"`
	SubjectCode string `json:"subject_code" binding:"required"`
}

type BulkClassSubjectsRequest struct {
	Items []ClassSubjectRequest `json:"items" binding:"required,dive"`
}