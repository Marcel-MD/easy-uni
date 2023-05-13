package services

import (
	"easy-uni/auth"
	"easy-uni/config"
	"easy-uni/repositories"
	"errors"
	"fmt"
	"sync"

	"easy-uni/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll() []models.User
	FindByID(id string) (models.User, error)
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

func (s *userService) FindAll() []models.User {
	log.Debug().Msg("Finding all users")

	return s.repository.FindAll()
}

func (s *userService) FindByID(id string) (models.User, error) {
	log.Debug().Str("id", id).Msg("Finding user")

	user, err := s.repository.FindByID(id)
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

	if user.Password == "" {
		user.Password = uuid.New().String()

		mail := models.Mail{
			To:      []string{user.Email},
			Subject: "Welcome to EasyUni",
			Body:    fmt.Sprintf("Welcome %s!\nYour password for EasyUni is <b>%s</b>", user.Name, user.Password),
		}

		go s.mailService.Send(mail)
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

	return auth.Generate(newUser.ID, s.cfg.TokenLifespan, s.cfg.ApiSecret)
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

	return auth.Generate(existingUser.ID, s.cfg.TokenLifespan, s.cfg.ApiSecret)
}

func (s *userService) verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
