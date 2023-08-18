package repositories

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(page, size int) ([]models.User, error)
	FindById(id string) (models.User, error)
	Create(t *models.User) error
	Update(t *models.User) error
	Delete(t *models.User) error

	FindByEmail(email string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
	data.Repository[models.User]
}

var (
	userOnce sync.Once
	userRepo UserRepository
)

func GetUserRepository() UserRepository {
	userOnce.Do(func() {
		log.Info().Msg("Initializing user repository")
		userRepo = &userRepository{
			db:         data.GetDB(),
			Repository: data.NewRepository[models.User](),
		}
	})
	return userRepo
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error

	return user, err
}
