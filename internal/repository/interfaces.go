package repository

import (
	"context"

	"revac_storage_app/internal/models"

	"github.com/google/uuid"
)

type ResumeRepository interface {

	// Создание резюме соискателя без контроля количества существующих
	CreateResume(ctx context.Context, resume *models.ResumePrimary) (string, error)

	// Создание резюме соискателя с контролем каличества уже существующих
	CreateResumeWithQntControl(ctx context.Context, MaxResumesPerUser int, resume *models.ResumePrimary) (string, error)

	// Получение резюме по идентификатору
	GetResumeById(ctx context.Context, resumeId uuid.UUID) (*models.ResumeInfo, error)

	// Получение списка резюме по идентификатору соискателя
	GetResumesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfResumesWithoutData, error)

	// Удаление резюме по идентификатору
	DeleteResumeById(ctx context.Context, resumeId uuid.UUID) error

	// Обновление резюме
	UpdateResume(ctx context.Context, resume *models.ResumeChange) error
}

type VacancyRepository interface {

	// Создание вакансии
	CreateVacancy(ctx context.Context, vacancy *models.VacancyPrimary) (string, error)

	// Получение вакансии её идентификатору
	GetVacancyById(ctx context.Context, id uuid.UUID) (*models.VacancyInfo, error)

	// Получение списка вакансий по идентификатору компании-работодателя
	GetVacanciesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfVacanciesWithoutData, error)

	// Удаление вакансии по её идентификатору
	DeleteVacancyById(ctx context.Context, id uuid.UUID) error

	// Обновление вакансии
	UpdateVacancy(ctx context.Context, vacancy *models.VacancyChange) error
}
