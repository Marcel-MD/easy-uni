package services

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data/repositories"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/rs/zerolog/log"
)

type EventService interface {
	FindAll(page, size int) ([]models.Event, error)
	FindById(id string) (models.Event, error)
	Create(event models.CreateEvent) (models.Event, error)
	Delete(id string) error
}

type eventService struct {
	repo repositories.EventRepository
}

var (
	eventOnce sync.Once
	eventSrv  EventService
)

func GetEventService() EventService {
	eventOnce.Do(func() {
		log.Info().Msg("Initializing event service")
		eventSrv = &eventService{
			repo: repositories.GetEventRepository(),
		}
	})
	return eventSrv
}

func (s *eventService) FindAll(page, size int) ([]models.Event, error) {
	return s.repo.FindAll(page, size)
}

func (s *eventService) FindById(id string) (models.Event, error) {
	return s.repo.FindById(id)
}

func (s *eventService) Create(event models.CreateEvent) (models.Event, error) {
	newEvent := models.Event{
		Name:       event.Name,
		URL:        event.URL,
		VisitorID:  event.VisitorID,
		CampaignID: event.CampaignID,
		Payload:    event.Payload,
		Meta:       event.Meta,
	}

	err := s.repo.Create(&newEvent)
	if err != nil {
		return models.Event{}, err
	}

	return newEvent, nil
}

func (s *eventService) Delete(id string) error {
	event, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&event)
}
