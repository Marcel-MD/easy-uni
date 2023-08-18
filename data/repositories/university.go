package repositories

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UniversityRepository interface {
	FindAll(page, size int) ([]models.University, error)
	FindById(id string) (models.University, error)
	Create(t *models.University) error
	Update(t *models.University) error
	Delete(t *models.University) error

	Find(name string, country string, city string) []models.University
}

type universityRepository struct {
	db *gorm.DB
	data.Repository[models.University]
}

var (
	universityOnce sync.Once
	universityRepo UniversityRepository
)

func GetUniversityRepository() UniversityRepository {
	universityOnce.Do(func() {
		log.Info().Msg("Initializing university repository")
		universityRepo = &universityRepository{
			db:         data.GetDB(),
			Repository: data.NewRepository[models.University](),
		}
	})
	return universityRepo
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
