package services

import (
	"github.com/Marcel-MD/easy-uni/models"
	"github.com/Marcel-MD/easy-uni/repositories"
)

type EventService interface {
	FindAll() []models.Event
	FindByID(id string) (models.Event, error)
	Create(event models.CreateEvent) (models.Event, error)
	Delete(id string) error
}

type eventService struct {
	repo repositories.EventRepository
}

func GetEventService() EventService {
	return &eventService{
		repo: repositories.GetEventRepository(),
	}
}

func (s *eventService) FindAll() []models.Event {
	return s.repo.FindAll()
}

func (s *eventService) FindByID(id string) (models.Event, error) {
	return s.repo.FindByID(id)
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
	event, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&event)
}
