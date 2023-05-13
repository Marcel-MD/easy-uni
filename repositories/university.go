package repositories

import (
	"easy-uni/models"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UniversityRepository interface {
	FindAll() []models.University
	FindByID(id string) (models.University, error)
	Find(name string, country string, city string) []models.University

	Create(university *models.University) error
	Update(university *models.University) error
	Delete(university *models.University) error
}

type universityRepository struct {
	db *gorm.DB
}

var (
	universityOnce sync.Once
	universityRepo UniversityRepository
)

func GetUniversityRepository() UniversityRepository {
	universityOnce.Do(func() {
		log.Info().Msg("Initializing university repository")
		universityRepo = &universityRepository{
			db: GetDB(),
		}
	})
	return universityRepo
}

func (r *universityRepository) FindAll() []models.University {
	var universities []models.University
	r.db.Find(&universities)
	return universities
}

func (r *universityRepository) FindByID(id string) (models.University, error) {
	var university models.University
	err := r.db.First(&university, "id = ?", id).Error

	return university, err
}

func (r *universityRepository) Find(name string, country string, city string) []models.University {
	var universities []models.University
	query := r.db.Model(&models.University{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if country != "" {
		query = query.Where("country LIKE ?", "%"+country+"%")
	}
	if city != "" {
		query = query.Where("city LIKE ?", "%"+city+"%")
	}
	query.Find(&universities)
	return universities
}

func (r *universityRepository) Create(university *models.University) error {
	return r.db.Create(university).Error
}

func (r *universityRepository) Update(university *models.University) error {
	return r.db.Save(university).Error
}

func (r *universityRepository) Delete(university *models.University) error {
	return r.db.Delete(university).Error
}
