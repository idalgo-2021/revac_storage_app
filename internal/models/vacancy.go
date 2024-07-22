package models

import (
	"time"

	"github.com/google/uuid"
)

// Для новых вакансий
type VacancyPrimary struct {
	OwnerId      string    `json:"owner_id"`
	CreateTime   time.Time `json:"create_time"`
	VacancyTitle string    `json:"vacancy_title"`
	DataContent  string    `json:"data_content"`
	IsActive     bool      `json:"is_active"`
	IsDraft      bool      `json:"is_draft"`
}

// Для изменения вакансии
type VacancyChange struct {
	ID           uuid.UUID `json:"id"`
	OwnerId      string    `json:"owner_id"`
	UpdateTime   time.Time `json:"update_time"`
	VacancyTitle string    `json:"resume_title"`
	DataContent  string    `json:"data_content"`
	IsActive     bool      `json:"is_active"`
	IsDraft      bool      `json:"is_draft"`
}

// Для передачи сновной(полной) информации по вакансии
type VacancyInfo struct {
	ID           uuid.UUID  `json:"id"`
	OwnerId      string     `json:"owner_id"`
	CreateTime   time.Time  `json:"create_time"`
	UpdateTime   *time.Time `json:"update_time"`
	VacancyTitle string     `json:"vacancy_title"`
	DataContent  string     `json:"data_content"`
	IsActive     bool       `json:"is_active"`
	IsDraft      bool       `json:"is_draft"`
}

// Для передачи только базовых атрибутов вакансии, без данных (для отображения в списках)
type VacancyWithoutData struct {
	ID           uuid.UUID  `json:"id"`
	OwnerId      string     `json:"owner_id"`
	CreateTime   time.Time  `json:"create_time"`
	UpdateTime   *time.Time `json:"update_time"`
	VacancyTitle string     `json:"vacancy_title"`
	IsActive     bool       `json:"is_active"`
	IsDraft      bool       `json:"is_draft"`
}

// Для передачи списка вакансий (состоящих только из базовых атрибутов)
type ListOfVacanciesWithoutData struct {
	Vacancies []VacancyWithoutData
}
