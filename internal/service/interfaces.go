package service

import (
	"context"
	"revac_storage_app/internal/models"

	"github.com/google/uuid"
)

type ResumeService interface {
	SCreateResume(ctx context.Context, resume *models.ResumePrimary) (string, error)
	SGetResumeById(ctx context.Context, id uuid.UUID) (*models.ResumeInfo, error)
	SGetResumesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfResumesWithoutData, error)
	SDeleteResumeById(ctx context.Context, id uuid.UUID) error
	SUpdateResume(ctx context.Context, resume *models.ResumeChange) error
}

type VacancyService interface {
	SCreateVacancy(ctx context.Context, vacancy *models.VacancyPrimary) (string, error)
	SGetVacancyById(ctx context.Context, id uuid.UUID) (*models.VacancyInfo, error)
	SGetVacanciesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfVacanciesWithoutData, error)
	SDeleteVacancyById(ctx context.Context, id uuid.UUID) error
	SUpdateVacancy(ctx context.Context, vacancy *models.VacancyChange) error
}
