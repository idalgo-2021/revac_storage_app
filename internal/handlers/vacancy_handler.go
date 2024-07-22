package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"revac_storage_app/internal/models"
	"revac_storage_app/internal/service"

	"github.com/google/uuid"
)

type VacancyHandler struct {
	service service.VacancyService
}

func NewVacancyHandler(s service.VacancyService) *VacancyHandler {
	return &VacancyHandler{service: s}
}

// Получить вакансию по ID
func (h *VacancyHandler) HandleGetVacancyById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid id format", http.StatusBadRequest)
		return
	}

	vacancy, err := h.service.SGetVacancyById(r.Context(), parsedId)
	if err != nil {
		if err == service.ErrVacancyNotFound {
			http.Error(w, "Vacancy not found", http.StatusNotFound)
		} else {
			log.Printf("Error getting vacancy: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vacancy); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Получить список вакансий с отбором по параметрам
func (h *VacancyHandler) HandleGetVacancies(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id != "" {
		h.HandleGetVacancyById(w, r)
		return
	}

	owner_id := r.URL.Query().Get("owner_id")
	if owner_id != "" {
		h.HandleGetVacanciesByOwnerId(w, r)
		return
	}

	// Больше ничего не реализовываем
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Получить список вакансий по owner_id
func (h *VacancyHandler) HandleGetVacanciesByOwnerId(w http.ResponseWriter, r *http.Request) {

	owner_id := r.URL.Query().Get("owner_id")
	if owner_id == "" {
		http.Error(w, "Missing owner_id parameter", http.StatusBadRequest)
		return
	}

	vacancies, err := h.service.SGetVacanciesByOwnerId(r.Context(), owner_id)
	if err != nil {
		if err == service.ErrVacancyNotFound {
			http.Error(w, "Vacancies not found", http.StatusNotFound)
		} else {
			log.Printf("Error getting vacancies: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vacancies); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Создать вакансию
func (h *VacancyHandler) HandleCreateVacancy(w http.ResponseWriter, r *http.Request) {

	var vacancy models.VacancyPrimary
	if err := json.NewDecoder(r.Body).Decode(&vacancy); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.service.SCreateVacancy(r.Context(), &vacancy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"id": id}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Удалить вакансию по ID
func (h *VacancyHandler) HandleDeleteVacancyById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid id format", http.StatusBadRequest)
		return
	}

	err = h.service.SDeleteVacancyById(r.Context(), parsedId)
	if err != nil {
		if err == service.ErrVacancyNotFound {
			http.Error(w, "Vacancy not found", http.StatusNotFound)
		} else {
			log.Printf("Error deleting vacancy: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *VacancyHandler) HandleUpdateVacancy(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid id format", http.StatusBadRequest)
		return
	}

	var vacancy models.VacancyChange
	if err := json.NewDecoder(r.Body).Decode(&vacancy); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	vacancy.ID = parsedId

	err = h.service.SUpdateVacancy(r.Context(), &vacancy)
	if err != nil {
		if errors.Is(err, service.ErrVacancyNotFound) {
			http.Error(w, "Vacancy not found", http.StatusNotFound)
		} else {
			log.Printf("Error updating vacancy: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
