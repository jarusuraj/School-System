package models

type Subjects struct {
}
type SubjectRequest struct {
	ID          int    `json:"subject_id"`
	SubjectName string `json:"subject_name" binding:"required"`
}

type BulkCreateSubjectsRequest struct {
	Subjects []SubjectRequest `json:"subjects" binding:"required,dive"`
}
