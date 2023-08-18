package services

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data/repositories"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/rs/zerolog/log"
)

type UniversityService interface {
	FindAll(page, size int) ([]models.University, error)
	FindById(id string) (models.University, error)
	Find(query models.UniversityQuery) []models.University

	Create(university models.CreateUniversity) (models.University, error)
	Update(id string, university models.CreateUniversity) (models.University, error)
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

func (s *universityService) FindAll(page, size int) ([]models.University, error) {
	return s.repo.FindAll(page, size)
}

func (s *universityService) FindById(id string) (models.University, error) {
	return s.repo.FindById(id)
}

func (s *universityService) Find(query models.UniversityQuery) []models.University {
	return s.repo.Find(query.Name, query.Country, query.City)
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

func (s *universityService) Update(id string, university models.CreateUniversity) (models.University, error) {
	universityToUpdate, err := s.repo.FindById(id)
	if err != nil {
		return models.University{}, err
	}

	universityToUpdate.Name = university.Name
	universityToUpdate.About = university.About
	universityToUpdate.Country = university.Country
	universityToUpdate.City = university.City
	universityToUpdate.Ranking = university.Ranking
	universityToUpdate.ImgLink = university.ImgLink

	err = s.repo.Update(&universityToUpdate)
	if err != nil {
		return models.University{}, err
	}

	return universityToUpdate, nil
}

func (s *universityService) Delete(id string) error {
	university, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&university)
}
