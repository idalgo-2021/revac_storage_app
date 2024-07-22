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

type ResumeHandler struct {
	service service.ResumeService
}

func NewResumeHandler(s service.ResumeService) *ResumeHandler {
	return &ResumeHandler{service: s}
}

// Получить резюме по ID
func (h *ResumeHandler) HandleGetResumeById(w http.ResponseWriter, r *http.Request) {

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

	resume, err := h.service.SGetResumeById(r.Context(), parsedId)
	if err != nil {
		if err == service.ErrResumeNotFound {
			http.Error(w, "Resume not found", http.StatusNotFound)
		} else {
			log.Printf("Error getting resume: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resume); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Создать резюме
func (h *ResumeHandler) HandleCreateResume(w http.ResponseWriter, r *http.Request) {

	var resume models.ResumePrimary
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.service.SCreateResume(r.Context(), &resume)
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

// Получить список резюме по owner_id
func (h *ResumeHandler) HandleGetResumesByOwnerId(w http.ResponseWriter, r *http.Request) {

	owner_id := r.URL.Query().Get("owner_id")
	if owner_id == "" {
		http.Error(w, "Missing owner_id parameter", http.StatusBadRequest)
		return
	}

	resumes, err := h.service.SGetResumesByOwnerId(r.Context(), owner_id)
	if err != nil {
		if err == service.ErrResumeNotFound {
			http.Error(w, "Resumes not found", http.StatusNotFound)
		} else {
			log.Printf("Error getting resumes: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resumes); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Получить список резюме с отбором по параметрам
func (h *ResumeHandler) HandleGetResumes(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id != "" {
		h.HandleGetResumeById(w, r)
		return
	}

	owner_id := r.URL.Query().Get("owner_id")
	if owner_id != "" {
		h.HandleGetResumesByOwnerId(w, r)
		return
	}

	// Больше ничего не реализовываем
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Удалить резюме по ID
func (h *ResumeHandler) HandleDeleteResumeById(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.SDeleteResumeById(r.Context(), parsedId)
	if err != nil {
		if err == service.ErrResumeNotFound {
			http.Error(w, "Resume not found", http.StatusNotFound)
		} else {
			log.Printf("Error deleting resume: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Обновить резюме
func (h *ResumeHandler) HandleUpdateResume(w http.ResponseWriter, r *http.Request) {
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

	var resume models.ResumeChange
	if err := json.NewDecoder(r.Body).Decode(&resume); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	resume.ID = parsedId

	err = h.service.SUpdateResume(r.Context(), &resume)
	if err != nil {
		if errors.Is(err, service.ErrResumeNotFound) {
			http.Error(w, "Resume not found", http.StatusNotFound)
		} else {
			log.Printf("Error updating resume: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
