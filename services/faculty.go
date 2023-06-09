package services

import (
	"github.com/Marcel-MD/easy-uni/models"
	"github.com/Marcel-MD/easy-uni/repositories"
)

type FacultyService interface {
	FindAll() []models.Faculty
	FindByID(id string) (models.Faculty, error)
	Find(name string, country string, city string, domain string, budget int) []models.Faculty

	Create(universityID string, faculty models.CreateFaculty) (models.Faculty, error)
	Update(id string, faculty models.CreateFaculty) (models.Faculty, error)
	Delete(id string) error
}

type facultyService struct {
	repo repositories.FacultyRepository
}

func GetFacultyService() FacultyService {
	return &facultyService{
		repo: repositories.GetFacultyRepository(),
	}
}

func (s *facultyService) FindAll() []models.Faculty {
	return s.repo.FindAll()
}

func (s *facultyService) FindByID(id string) (models.Faculty, error) {
	return s.repo.FindByID(id)
}

func (s *facultyService) Find(name string, country string, city string, domain string, budget int) []models.Faculty {
	return s.repo.Find(name, country, city, domain, budget)
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
	facultyToUpdate, err := s.repo.FindByID(id)
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
	faculty, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&faculty)
}
