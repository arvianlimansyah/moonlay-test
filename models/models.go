package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required"`
	Attachment  string `gorm:"not null" json:"attachment"`
}

type Subtask struct {
	gorm.Model
	TaskId      int32  `gorm:"not null" json:"task_id" binding:"required"`
	Title       string `gorm:"not null" json:"title" binding:"required"`
	Description string `gorm:"not null" json:"description" binding:"required"`
	Attachment  string `gorm:"not null" json:"attachment"`
}

type TaskWithSubtask struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Attachment  string    `json:"attachment"`
	Subtasks    []Subtask `json:"subtasks"`
}

type HealthCheckResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
