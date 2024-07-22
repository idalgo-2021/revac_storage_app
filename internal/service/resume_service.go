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

var ErrResumeNotFound = errors.New("resume not found")

type resumeService struct {
	repo repository.ResumeRepository
	cfg  config.Config
}

func NewResumeService(repo repository.ResumeRepository, cfg config.Config) ResumeService {
	return &resumeService{repo: repo, cfg: cfg}
}

func (s *resumeService) SCreateResume(ctx context.Context, resume *models.ResumePrimary) (string, error) {

	// Проверка общих параметров на количество резюме
	ResumeServiceConfig := s.cfg.ResumeServiceConfig
	if ResumeServiceConfig.ControlQntResumesPerUserEnabled {
		if ResumeServiceConfig.MaxResumesPerUser <= 0 {
			return "", fmt.Errorf("configuration parameter: MaxResumesPerUser <= 0")
		}
	}

	// Проверка наличия данных
	if resume.OwnerId == "" || resume.DataContent == "" || resume.ResumeTitle == "" {
		return "", fmt.Errorf("missing required fields")
	}

	// Значения по умолчанию
	resume.CreateTime = time.Now().UTC()
	resume.IsActive = true
	resume.IsDraft = true

	var id string
	var err error
	if ResumeServiceConfig.ControlQntResumesPerUserEnabled {
		id, err = s.repo.CreateResumeWithQntControl(ctx, ResumeServiceConfig.MaxResumesPerUser, resume)
	} else {
		id, err = s.repo.CreateResume(ctx, resume)
	}
	if err != nil {
		return "", fmt.Errorf("failed to create resume: %w", err)
	}

	return id, nil
}

func (s *resumeService) SGetResumeById(ctx context.Context, id uuid.UUID) (*models.ResumeInfo, error) {

	resume, err := s.repo.GetResumeById(ctx, id)
	if err != nil {
		if err == customErrors.ErrNotFound {
			return nil, ErrResumeNotFound
		}
		return nil, err
	}
	return resume, nil
}

func (s *resumeService) SGetResumesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfResumesWithoutData, error) {

	resumes, err := s.repo.GetResumesByOwnerId(ctx, ownerId)
	if err != nil {
		if err == customErrors.ErrNotFound {
			return nil, ErrResumeNotFound
		}
		return nil, err
	}
	return resumes, nil
}

func (s *resumeService) SDeleteResumeById(ctx context.Context, id uuid.UUID) error {

	err := s.repo.DeleteResumeById(ctx, id)
	if err != nil {
		if err == customErrors.ErrNotFound {
			return ErrResumeNotFound
		}
		return err
	}
	return nil
}

func (s *resumeService) SUpdateResume(ctx context.Context, resume *models.ResumeChange) error {
	err := s.repo.UpdateResume(ctx, resume)
	if err != nil {
		if errors.Is(err, customErrors.ErrNotFound) {
			return fmt.Errorf("resume not found: %w", err)
		}
		return fmt.Errorf("failed to update resume: %w", err)
	}
	return nil
}
