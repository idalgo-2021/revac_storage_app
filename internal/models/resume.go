package models

import (
	"time"

	"github.com/google/uuid"
)

// Для новых резюме
type ResumePrimary struct {
	OwnerId     string    `json:"owner_id"`
	CreateTime  time.Time `json:"create_time"`
	ResumeTitle string    `json:"resume_title"`
	DataContent string    `json:"data_content"`
	IsActive    bool      `json:"is_active"`
	IsDraft     bool      `json:"is_draft"`
}

// Для изменения резюме
type ResumeChange struct {
	ID          uuid.UUID `json:"id"`
	OwnerId     string    `json:"owner_id"`
	UpdateTime  time.Time `json:"update_time"`
	ResumeTitle string    `json:"resume_title"`
	DataContent string    `json:"data_content"`
	IsActive    bool      `json:"is_active"`
	IsDraft     bool      `json:"is_draft"`
}

// Для передачи сновной(полной) информации из резюме
type ResumeInfo struct {
	ID          uuid.UUID  `json:"id"`
	OwnerId     string     `json:"owner_id"`
	CreateTime  time.Time  `json:"create_time"`
	UpdateTime  *time.Time `json:"update_time"`
	ResumeTitle string     `json:"resume_title"`
	DataContent string     `json:"data_content"`
	IsActive    bool       `json:"is_active"`
	IsDraft     bool       `json:"is_draft"`
}

// Для передачи только базовых атрибутов резюме, без данных (для отображения в списках)
type ResumeWithoutData struct {
	ID          uuid.UUID  `json:"id"`
	OwnerId     string     `json:"owner_id"`
	CreateTime  time.Time  `json:"create_time"`
	UpdateTime  *time.Time `json:"update_time"`
	ResumeTitle string     `json:"resume_title"`
	IsActive    bool       `json:"is_active"`
	IsDraft     bool       `json:"is_draft"`
}

// Для передачи списка резюме (состоящих только из базовых атрибутов)
type ListOfResumesWithoutData struct {
	Resumes []ResumeWithoutData
}
