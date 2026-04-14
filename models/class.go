package models
type CreateClassRequest struct {
	ClassNo int    `json:"class_no" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type BulkCreateClassRequest struct {
	Classes []CreateClassRequest `json:"classes" binding:"required,dive"`
}