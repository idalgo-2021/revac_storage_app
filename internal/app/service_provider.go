package app

import (
	"database/sql"
	"log"

	"revac_storage_app/internal/config"
	"revac_storage_app/internal/handlers"
	"revac_storage_app/internal/repository"
	"revac_storage_app/internal/service"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	resumeRepository  repository.ResumeRepository
	vacancyRepository repository.VacancyRepository

	resumeService  service.ResumeService
	vacancyService service.VacancyService

	resumeHandler  *handlers.ResumeHandler
	vacancyHandler *handlers.VacancyHandler
}

func newServiceProvider(db *sql.DB, cfg config.Config) *serviceProvider {
	resumeRepo := repository.NewResumeRepository(db)
	vacancyRepo := repository.NewVacancyRepository(db)

	resumeServ := service.NewResumeService(resumeRepo, cfg)
	vacancyServ := service.NewVacancyService(vacancyRepo, cfg)

	resumeHandler := handlers.NewResumeHandler(resumeServ)
	vacancyHandler := handlers.NewVacancyHandler(vacancyServ)

	return &serviceProvider{
		resumeRepository:  resumeRepo,
		vacancyRepository: vacancyRepo,

		resumeService:  resumeServ,
		vacancyService: vacancyServ,

		resumeHandler:  resumeHandler,
		vacancyHandler: vacancyHandler,
	}
}

func (s *serviceProvider) HTTPConfig(cfg *config.HTTPServerConfig) config.HTTPConfig {
	if s.httpConfig == nil {
		currHTTPConfig, err := config.NewHTTPConfig(cfg)
		if err != nil {
			log.Fatalf("failed to get http config params: %s", err.Error())
		}

		s.httpConfig = currHTTPConfig
	}

	return s.httpConfig
}
