package services

import (
	"easy-uni/auth"
	"easy-uni/config"
	"easy-uni/repositories"
	"errors"
	"sync"

	"easy-uni/models"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll() []models.User
	FindOne(id string) (models.User, error)
	Register(user models.RegisterUser) (string, error)
	Login(user models.LoginUser) (string, error)
}

type userService struct {
	repository repositories.UserRepository
	cfg        config.Config
}

var (
	userOnce sync.Once
	userSrv  UserService
)

func GetUserService() UserService {
	userOnce.Do(func() {
		log.Info().Msg("Initializing user service")

		userSrv = &userService{
			repository: repositories.GetUserRepository(),
			cfg:        config.GetConfig(),
		}
	})

	return userSrv
}

func (s *userService) FindAll() []models.User {
	log.Debug().Msg("Finding all users")

	return s.repository.FindAll()
}

func (s *userService) FindOne(id string) (models.User, error) {
	log.Debug().Str("id", id).Msg("Finding user")

	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) Register(user models.RegisterUser) (string, error) {
	log.Debug().Msg("Registering user")

	_, err := s.repository.FindByEmail(user.Email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := models.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: string(hashedPassword),
	}

	err = s.repository.Create(&newUser)
	if err != nil {
		return "", err
	}

	return auth.Generate(newUser.Id, s.cfg.TokenLifespan, s.cfg.ApiSecret)
}

func (s *userService) Login(user models.LoginUser) (string, error) {
	log.Debug().Msg("Logging in user")

	existingUser, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = s.verifyPassword(user.Password, existingUser.Password)
	if err != nil {
		return "", err
	}

	return auth.Generate(existingUser.Id, s.cfg.TokenLifespan, s.cfg.ApiSecret)
}

func (s *userService) verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
