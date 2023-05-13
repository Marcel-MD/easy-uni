package repositories

import (
	"easy-uni/models"
	"sync"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type FacultyRepository interface {
	FindAll() []models.Faculty
	FindByID(id string) (models.Faculty, error)
	FindByUniversityID(universityID string) []models.Faculty
	Find(name string, country string, city string, domain string, budget int) []models.Faculty

	Create(faculty *models.Faculty) error
	Update(faculty *models.Faculty) error
	Delete(faculty *models.Faculty) error
}

type facultyRepository struct {
	db *gorm.DB
}

var (
	facultyOnce sync.Once
	facultyRepo FacultyRepository
)

func GetFacultyRepository() FacultyRepository {
	facultyOnce.Do(func() {
		log.Info().Msg("Initializing faculty repository")
		facultyRepo = &facultyRepository{
			db: GetDB(),
		}
	})
	return facultyRepo
}

func (r *facultyRepository) FindAll() []models.Faculty {
	var faculties []models.Faculty
	r.db.Find(&faculties)
	return faculties
}

func (r *facultyRepository) FindByID(id string) (models.Faculty, error) {
	var faculty models.Faculty
	err := r.db.First(&faculty, "id = ?", id).Error

	return faculty, err
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

func (r *facultyRepository) Create(faculty *models.Faculty) error {
	return r.db.Create(&faculty).Error
}

func (r *facultyRepository) Update(faculty *models.Faculty) error {
	return r.db.Save(&faculty).Error
}

func (r *facultyRepository) Delete(faculty *models.Faculty) error {
	return r.db.Delete(&faculty).Error
}
