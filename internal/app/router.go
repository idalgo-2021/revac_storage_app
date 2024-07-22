package app

import (
	"net/http"
	"strings"
)

func (s *serviceProvider) Router() http.Handler {
	mux := http.NewServeMux()

	// Handler for specific resume operations (GET, PUT, DELETE)
	mux.HandleFunc("/resume/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/resume/")
		if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		query := r.URL.Query()
		query.Set("id", id)
		r.URL.RawQuery = query.Encode()

		switch r.Method {
		case http.MethodGet:
			s.resumeHandler.HandleGetResumeById(w, r)
		case http.MethodPut:
			s.resumeHandler.HandleUpdateResume(w, r)
		case http.MethodDelete:
			s.resumeHandler.HandleDeleteResumeById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handler for creating and listing resumes
	mux.HandleFunc("/resumes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.resumeHandler.HandleGetResumes(w, r)
		case http.MethodPost:
			s.resumeHandler.HandleCreateResume(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handler for specific vacancy operations (GET, PUT, DELETE)
	mux.HandleFunc("/vacancy/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/vacancy/")
		if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		query := r.URL.Query()
		query.Set("id", id)
		r.URL.RawQuery = query.Encode()

		switch r.Method {
		case http.MethodGet:
			s.vacancyHandler.HandleGetVacancyById(w, r)
		case http.MethodPut:
			s.vacancyHandler.HandleUpdateVacancy(w, r)
		case http.MethodDelete:
			s.vacancyHandler.HandleDeleteVacancyById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handler for creating and listing vacancies
	mux.HandleFunc("/vacancies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.vacancyHandler.HandleGetVacancies(w, r)
		case http.MethodPost:
			s.vacancyHandler.HandleCreateVacancy(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
