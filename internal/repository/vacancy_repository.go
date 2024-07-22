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

type vacancyRepository struct {
	db *sql.DB
}

func NewVacancyRepository(db *sql.DB) VacancyRepository {
	return &vacancyRepository{db: db}
}

func (r *vacancyRepository) CreateVacancy(ctx context.Context, vacancy *models.VacancyPrimary) (string, error) {

	query := `
		INSERT INTO vacancies (owner_id, create_time, vacancy_title, data_content)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id string
	err := r.db.QueryRowContext(ctx, query, vacancy.OwnerId, vacancy.CreateTime, vacancy.VacancyTitle, vacancy.DataContent).Scan(&id)
	if err != nil {
		log.Printf("Failed to execute query: %v", err)
		return "", err
	}
	return id, nil
}

func (r *vacancyRepository) GetVacancyById(ctx context.Context, id uuid.UUID) (*models.VacancyInfo, error) {

	query := `SELECT id, owner_id, create_time, update_time, vacancy_title, data_content FROM vacancies WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	var vacancy models.VacancyInfo
	err := row.Scan(&vacancy.ID, &vacancy.OwnerId, &vacancy.CreateTime, &vacancy.UpdateTime, &vacancy.VacancyTitle, &vacancy.DataContent)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get vacancy: %w", err)
	}

	return &vacancy, nil
}

func (r *vacancyRepository) GetVacanciesByOwnerId(ctx context.Context, ownerId string) (*models.ListOfVacanciesWithoutData, error) {

	query := "SELECT id, owner_id, create_time, update_time, vacancy_title FROM vacancies WHERE owner_id = $1"
	rows, err := r.db.QueryContext(ctx, query, ownerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []models.VacancyWithoutData
	for rows.Next() {
		var vacancy models.VacancyWithoutData
		if err := rows.Scan(&vacancy.ID, &vacancy.OwnerId, &vacancy.CreateTime, &vacancy.UpdateTime, &vacancy.VacancyTitle); err != nil {
			return nil, err
		}
		vacancies = append(vacancies, vacancy)
	}

	return &models.ListOfVacanciesWithoutData{Vacancies: vacancies}, nil
}

func (r *vacancyRepository) DeleteVacancyById(ctx context.Context, id uuid.UUID) error {

	query := "DELETE FROM vacancies WHERE id = $1"
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

func (r *vacancyRepository) UpdateVacancy(ctx context.Context, vacancy *models.VacancyChange) error {

	query := `UPDATE vacancies SET owner_id = $1, update_time = $2, vacancy_title = $3, data_content = $4 WHERE id = $5`
	result, err := r.db.ExecContext(ctx, query, vacancy.OwnerId, vacancy.UpdateTime, vacancy.VacancyTitle, vacancy.DataContent, vacancy.ID)
	if err != nil {
		log.Printf("Failed to update vacancy: %v", err)
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
