package services

import (
	"easy-uni/models"
	"easy-uni/repositories"
	"sync"

	"github.com/rs/zerolog/log"
)

type UniversityService interface {
	FindAll() []models.University
	FindByID(id string) (models.University, error)
	Find(name string, country string, city string) []models.University

	Create(university models.CreateUniversity) (models.University, error)
	Delete(id string) error
}

type universityService struct {
	repo repositories.UniversityRepository
}

var (
	universityOnce sync.Once
	universitySrv  UniversityService
)

func GetUniversityService() UniversityService {
	universityOnce.Do(func() {
		log.Info().Msg("Initializing university service")
		universitySrv = &universityService{
			repo: repositories.GetUniversityRepository(),
		}
	})
	return universitySrv
}

func (s *universityService) FindAll() []models.University {
	return s.repo.FindAll()
}

func (s *universityService) FindByID(id string) (models.University, error) {
	return s.repo.FindByID(id)
}

func (s *universityService) Find(name string, country string, city string) []models.University {
	return s.repo.Find(name, country, city)
}

func (s *universityService) Create(university models.CreateUniversity) (models.University, error) {
	newUniversity := models.University{
		Name:    university.Name,
		About:   university.About,
		Country: university.Country,
		City:    university.City,
		Ranking: university.Ranking,
		ImgLink: university.ImgLink,
	}

	err := s.repo.Create(&newUniversity)
	if err != nil {
		return models.University{}, err
	}

	return newUniversity, nil
}

func (s *universityService) Delete(id string) error {
	university, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&university)
}
