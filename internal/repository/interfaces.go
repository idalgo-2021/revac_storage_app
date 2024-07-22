package repository

import (
	"context"

	"revac_storage_app/internal/models"

	"github.com/google/uuid"
)

type ResumeRepository interface {

	// Создание резюме соискателя без контроля количества существующих
	CreateResume(ctx context.Context, infoData *models.ResumePrimary) (string, error)

	// Создание резюме соискателя с контролем каличества уже существующих
	CreateResumeWithQntControl(ctx context.Context, MaxResumesPerUser int, infoData *models.ResumePrimary) (string, error)

	// Получение резюме по его идентификатору
	GetResumeById(ctx context.Context, resumeId uuid.UUID) (*models.ResumeInfo, error)

	// Получение списка резюме по идентификатору соискателя
	GetResumesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfResumesWithoutData, error)

	// Удаление резюме по его идентификатору
	DeleteResumeById(ctx context.Context, resumeId uuid.UUID) error

	// Обновление резюме
	// UpdateResumeById(ctx context.Context, infoData *models.ResumeChange) error
}

type VacancyRepository interface {

	// Создание вакансии
	CreateVacancy(ctx context.Context, infoData *models.VacancyPrimary) (string, error)

	// Получение вакансии её идентификатору
	GetVacancyById(ctx context.Context, id uuid.UUID) (*models.VacancyInfo, error)

	// Получение списка вакансий по идентификатору компании-работодателя
	GetVacanciesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfVacanciesWithoutData, error)

	// Удаление вакансии по его идентификатору
	DeleteVacancyById(ctx context.Context, id uuid.UUID) error

	// Обновление вакансии
	// UpdateVacancyById(ctx context.Context, id uuid.UUID,) error
}
