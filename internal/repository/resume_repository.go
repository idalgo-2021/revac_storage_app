package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"revac_storage_app/internal/models"
	customErrors "revac_storage_app/internal/repository/errors"

	"github.com/google/uuid"
)

type resumeRepository struct {
	db *sql.DB
}

func NewResumeRepository(db *sql.DB) ResumeRepository {
	return &resumeRepository{db: db}
}

func (r *resumeRepository) CreateResume(ctx context.Context, resume *models.ResumePrimary) (string, error) {

	query := `
		INSERT INTO resumes (owner_id, create_time, resume_title, data_content)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id string
	err := r.db.QueryRowContext(ctx, query, resume.OwnerId, resume.CreateTime, resume.ResumeTitle, resume.DataContent).Scan(&id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return "", err
	}
	return id, nil
}

func (r *resumeRepository) CreateResumeWithQntControl(ctx context.Context, MaxResumesPerUser int, resume *models.ResumePrimary) (string, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return "", err
	}

	query := `
	WITH resume_count AS (
		SELECT COUNT(*) AS cnt
		FROM resumes
		WHERE owner_id = $1
	)
	INSERT INTO resumes (owner_id, create_time, resume_title, data_content)
	SELECT $1, $2, $3, $4
	FROM resume_count
	WHERE resume_count.cnt < $5
	RETURNING id;`

	row := tx.QueryRowContext(ctx, query, resume.OwnerId, resume.CreateTime, resume.ResumeTitle, resume.DataContent, MaxResumesPerUser)
	var id string
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			tx.Rollback()
			return "", fmt.Errorf("cannot create resume: limit of %d resumes per user reached", MaxResumesPerUser)
		}
		tx.Rollback()
		log.Printf("Failed to execute query: %v", err)
		return "", err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return "", err
	}

	return id, nil
}

func (r *resumeRepository) GetResumeById(ctx context.Context, id uuid.UUID) (*models.ResumeInfo, error) {

	query := `SELECT id, owner_id, create_time, update_time, resume_title, data_content FROM resumes WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var resume models.ResumeInfo
	err := row.Scan(&resume.ID, &resume.OwnerId, &resume.CreateTime, &resume.UpdateTime, &resume.ResumeTitle, &resume.DataContent)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get resume: %w", err)
	}

	return &resume, nil
}

func (r *resumeRepository) GetResumesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfResumesWithoutData, error) {

	query := "SELECT id, owner_id, create_time, update_time, resume_title FROM resumes WHERE owner_id = $1"
	rows, err := r.db.QueryContext(ctx, query, ownerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []models.ResumeWithoutData
	for rows.Next() {
		var resume models.ResumeWithoutData
		if err := rows.Scan(&resume.ID, &resume.OwnerId, &resume.CreateTime, &resume.UpdateTime, &resume.ResumeTitle); err != nil {
			return nil, err
		}
		resumes = append(resumes, resume)
	}

	return &models.ListOfResumesWithoutData{Resumes: resumes}, nil
}

func (r *resumeRepository) DeleteResumeById(ctx context.Context, id uuid.UUID) error {

	query := "DELETE FROM resumes WHERE id = $1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return customErrors.ErrNotFound
	}

	return nil
}

func (r *resumeRepository) UpdateResume(ctx context.Context, resume *models.ResumeChange) error {

	query := `UPDATE resumes SET owner_id = $1, update_time = $2, resume_title = $3, data_content = $4 WHERE id = $5`
	result, err := r.db.ExecContext(ctx, query, resume.OwnerId, resume.UpdateTime, resume.ResumeTitle, resume.DataContent, resume.ID)
	if err != nil {
		log.Printf("Failed to update resume: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return customErrors.ErrNotFound
	}
	return nil
}
