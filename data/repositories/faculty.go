package repositories

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type FacultyRepository interface {
	FindAll(page, size int) ([]models.Faculty, error)
	FindById(id string) (models.Faculty, error)
	Create(t *models.Faculty) error
	Update(t *models.Faculty) error
	Delete(t *models.Faculty) error

	FindByUniversityID(universityID string) []models.Faculty
	Find(name string, country string, city string, domain string, budget int) []models.Faculty
}

type facultyRepository struct {
	db *gorm.DB
	data.Repository[models.Faculty]
}

var (
	facultyOnce sync.Once
	facultyRepo FacultyRepository
)

func GetFacultyRepository() FacultyRepository {
	facultyOnce.Do(func() {
		log.Info().Msg("Initializing faculty repository")
		facultyRepo = &facultyRepository{
			db:         data.GetDB(),
			Repository: data.NewRepository[models.Faculty](),
		}
	})
	return facultyRepo
}

func (r *facultyRepository) FindByUniversityID(universityID string) []models.Faculty {
	var faculties []models.Faculty
	r.db.Where("university_id = ?", universityID).Find(&faculties)
	return faculties
}

func (r *facultyRepository) Find(name string, country string, city string, domain string, budget int) []models.Faculty {
	var faculties []models.Faculty
	query := r.db.Model(&models.Faculty{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if domain != "" {
		query = query.Where("domains @> ?", pq.Array([]string{domain}))
	}
	if budget > 0 {
		query = query.Where("budget <= ?", budget)
	}

	query = query.Joins("JOIN universities ON universities.id = faculties.university_id")

	if country != "" {
		query = query.Where("universities.country LIKE ?", "%"+country+"%")
	}
	if city != "" {
		query = query.Where("universities.city LIKE ?", "%"+city+"%")
	}

	query.Preload("University").Find(&faculties)
	return faculties
}
