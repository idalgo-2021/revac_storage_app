package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"revac_storage_app/internal/config"
	"revac_storage_app/internal/models"
	"revac_storage_app/internal/repository"
	customErrors "revac_storage_app/internal/repository/errors"

	"github.com/google/uuid"
)

var ErrVacancyNotFound = errors.New("vacancy not found")

type vacancyService struct {
	repo repository.VacancyRepository
	cfg  config.Config
}

func NewVacancyService(repo repository.VacancyRepository, cfg config.Config) VacancyService {
	return &vacancyService{repo: repo, cfg: cfg}
}

func (s *vacancyService) SCreateVacancy(ctx context.Context, infoData *models.VacancyPrimary) (string, error) {

	// Проверка наличия данных
	if infoData.OwnerId == "" || infoData.DataContent == "" || infoData.VacancyTitle == "" {
		return "", fmt.Errorf("missing required fields")
	}

	// Значения по умолчанию
	infoData.CreateTime = time.Now().UTC()
	infoData.IsActive = true
	infoData.IsDraft = true

	id, err := s.repo.CreateVacancy(ctx, infoData)
	if err != nil {
		return "", fmt.Errorf("failed to create vacancy: %w", err)
	}

	return id, nil
}

func (s *vacancyService) SGetVacancyById(ctx context.Context, id uuid.UUID) (*models.VacancyInfo, error) {

	vacancy, err := s.repo.GetVacancyById(ctx, id)
	if err != nil {
		if err == customErrors.ErrNotFound {
			return nil, ErrVacancyNotFound
		}
		return nil, err
	}
	return vacancy, nil
}

func (s *vacancyService) SGetVacanciesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfVacanciesWithoutData, error) {

	vacancies, err := s.repo.GetVacanciesByOwnerId(ctx, ownerId)
	if err != nil {
		if err == customErrors.ErrNotFound {
			return nil, ErrResumeNotFound
		}
		return nil, err
	}
	return vacancies, nil
}

func (s *vacancyService) SDeleteVacancyById(ctx context.Context, id uuid.UUID) error {

	err := s.repo.DeleteVacancyById(ctx, id)
	if err != nil {
		if err == customErrors.ErrNotFound {
			return ErrVacancyNotFound
		}
		return err
	}
	return nil

}
