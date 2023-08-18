package services

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data/repositories"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/rs/zerolog/log"
)

type FacultyService interface {
	FindAll(page, size int) ([]models.Faculty, error)
	FindById(id string) (models.Faculty, error)
	Find(query models.FacultyQuery) []models.Faculty

	Create(universityID string, faculty models.CreateFaculty) (models.Faculty, error)
	Update(id string, faculty models.CreateFaculty) (models.Faculty, error)
	Delete(id string) error
}

type facultyService struct {
	repo repositories.FacultyRepository
}

var (
	facultyOnce sync.Once
	facultySrv  FacultyService
)

func GetFacultyService() FacultyService {
	facultyOnce.Do(func() {
		log.Info().Msg("Initializing faculty service")
		facultySrv = &facultyService{
			repo: repositories.GetFacultyRepository(),
		}
	})
	return facultySrv
}

func (s *facultyService) FindAll(page, size int) ([]models.Faculty, error) {
	return s.repo.FindAll(page, size)
}

func (s *facultyService) FindById(id string) (models.Faculty, error) {
	return s.repo.FindById(id)
}

func (s *facultyService) Find(query models.FacultyQuery) []models.Faculty {
	return s.repo.Find(query.Name, query.Country, query.City, query.Domain, query.Budget)
}

func (s *facultyService) Create(universityID string, faculty models.CreateFaculty) (models.Faculty, error) {
	newFaculty := models.Faculty{
		Name:                 faculty.Name,
		Domains:              faculty.Domains,
		About:                faculty.About,
		Budget:               faculty.Budget,
		Duration:             faculty.Duration,
		ApplyDate:            faculty.ApplyDate,
		StartDate:            faculty.StartDate,
		AcademicRequirements: faculty.AcademicRequirements,
		LanguageRequirements: faculty.LanguageRequirements,
		OtherRequirements:    faculty.OtherRequirements,
		UniversityID:         universityID,
	}

	err := s.repo.Create(&newFaculty)
	if err != nil {
		return models.Faculty{}, err
	}

	return newFaculty, nil
}

func (s *facultyService) Update(id string, faculty models.CreateFaculty) (models.Faculty, error) {
	facultyToUpdate, err := s.repo.FindById(id)
	if err != nil {
		return models.Faculty{}, err
	}

	facultyToUpdate.Name = faculty.Name
	facultyToUpdate.Domains = faculty.Domains
	facultyToUpdate.About = faculty.About
	facultyToUpdate.Budget = faculty.Budget
	facultyToUpdate.Duration = faculty.Duration
	facultyToUpdate.ApplyDate = faculty.ApplyDate
	facultyToUpdate.StartDate = faculty.StartDate
	facultyToUpdate.AcademicRequirements = faculty.AcademicRequirements
	facultyToUpdate.LanguageRequirements = faculty.LanguageRequirements
	facultyToUpdate.OtherRequirements = faculty.OtherRequirements

	err = s.repo.Update(&facultyToUpdate)
	if err != nil {
		return models.Faculty{}, err
	}

	return facultyToUpdate, nil
}

func (s *facultyService) Delete(id string) error {
	faculty, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&faculty)
}
