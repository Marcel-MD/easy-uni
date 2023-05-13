package repositories

import (
	"easy-uni/models"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []models.User
	FindByID(id string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

var (
	userOnce sync.Once
	userRepo UserRepository
)

func GetUserRepository() UserRepository {
	userOnce.Do(func() {
		log.Info().Msg("Initializing user repository")
		userRepo = &userRepository{
			db: GetDB(),
		}
	})
	return userRepo
}

func (r *userRepository) FindAll() []models.User {
	var users []models.User
	r.db.Find(&users)
	return users
}

func (r *userRepository) FindByID(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error

	return user, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error

	return user, err
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}
