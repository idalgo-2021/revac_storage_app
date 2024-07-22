package service

import (
	"context"
	"revac_storage_app/internal/models"

	"github.com/google/uuid"
)

type ResumeService interface {
	SCreateResume(ctx context.Context, infoData *models.ResumePrimary) (string, error)
	SGetResumeById(ctx context.Context, id uuid.UUID) (*models.ResumeInfo, error)
	SGetResumesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfResumesWithoutData, error)
	SDeleteResumeById(ctx context.Context, id uuid.UUID) error
	// SUpdateResumeById(ctx context.Context, id uuid.UUID) error

}

type VacancyService interface {
	SCreateVacancy(ctx context.Context, infoData *models.VacancyPrimary) (string, error)
	SGetVacancyById(ctx context.Context, id uuid.UUID) (*models.VacancyInfo, error)
	SGetVacanciesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfVacanciesWithoutData, error)
	SDeleteVacancyById(ctx context.Context, id uuid.UUID) error
	// SUpdateVacancyById(ctx context.Context, id uuid.UUID) error
}
