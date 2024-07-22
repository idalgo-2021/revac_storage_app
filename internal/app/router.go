package app

import (
	"net/http"
	"strings"
)

func (s *serviceProvider) Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/resume/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/resume/")
		if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		query := r.URL.Query()
		query.Set("id", id)
		r.URL.RawQuery = query.Encode()
		s.resumeHandler.HandleGetResumeById(w, r)
	})

	mux.HandleFunc("/resumes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.resumeHandler.HandleGetResumes(w, r)
		case http.MethodPost:
			s.resumeHandler.HandleCreateResume(w, r)
		case http.MethodDelete:
			s.resumeHandler.HandleDeleteResumeById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/vacancy/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/vacancy/")
		if id == "" {
			http.Error(w, "Missing id parameter", http.StatusBadRequest)
			return
		}
		query := r.URL.Query()
		query.Set("id", id)
		r.URL.RawQuery = query.Encode()
		s.vacancyHandler.HandleGetVacancyById(w, r)
	})

	mux.HandleFunc("/vacancies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.vacancyHandler.HandleGetVacancies(w, r)
		case http.MethodPost:
			s.vacancyHandler.HandleCreateVacancy(w, r)
		case http.MethodDelete:
			s.vacancyHandler.HandleDeleteVacancyById(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// mux.HandleFunc("/employer", s.employerHandler.HandleRequest)

	return mux
}
