package services

import (
	"errors"
	"sync"

	"github.com/Marcel-MD/easy-uni/auth"
	"github.com/Marcel-MD/easy-uni/config"
	"github.com/Marcel-MD/easy-uni/data/repositories"

	"github.com/Marcel-MD/easy-uni/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll(page, size int) ([]models.User, error)
	FindById(id string) (models.User, error)
	Register(user models.RegisterUser) (string, error)
	Login(user models.LoginUser) (string, error)
}

type userService struct {
	mailService MailService
	repository  repositories.UserRepository
	cfg         config.Config
}

var (
	userOnce sync.Once
	userSrv  UserService
)

func GetUserService() UserService {
	userOnce.Do(func() {
		log.Info().Msg("Initializing user service")
		userSrv = &userService{
			mailService: GetMailService(),
			repository:  repositories.GetUserRepository(),
			cfg:         config.GetConfig(),
		}
	})
	return userSrv
}

func (s *userService) FindAll(page, size int) ([]models.User, error) {
	return s.repository.FindAll(page, size)
}

func (s *userService) FindById(id string) (models.User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) Register(user models.RegisterUser) (string, error) {
	_, err := s.repository.FindByEmail(user.Email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	if user.Password == "" {
		user.Password = uuid.New().String()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	newUser := models.User{
		Email:     user.Email,
		Name:      user.Name,
		Password:  string(hashedPassword),
		VisitorID: user.VisitorID,
		Roles:     []string{models.UserRole},
	}

	err = s.repository.Create(&newUser)
	if err != nil {
		return "", err
	}

	return auth.Generate(newUser.ID, newUser.Roles, s.cfg.TokenLifespan, s.cfg.ApiSecret)
}

func (s *userService) Login(user models.LoginUser) (string, error) {
	existingUser, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = s.verifyPassword(user.Password, existingUser.Password)
	if err != nil {
		return "", err
	}

	return auth.Generate(existingUser.ID, existingUser.Roles, s.cfg.TokenLifespan, s.cfg.ApiSecret)
}

func (s *userService) verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
