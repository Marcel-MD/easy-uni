package repositories

import (
	"easy-uni/models"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll() []models.Event
	FindByID(id string) (models.Event, error)
	Create(event *models.Event) error
	Delete(event *models.Event) error
}

type eventRepository struct {
	db *gorm.DB
}

var (
	eventOnce sync.Once
	eventRepo EventRepository
)

func GetEventRepository() EventRepository {
	eventOnce.Do(func() {
		log.Info().Msg("Initializing event repository")
		eventRepo = &eventRepository{
			db: GetDB(),
		}
	})
	return eventRepo
}

func (r *eventRepository) FindAll() []models.Event {
	var events []models.Event
	r.db.Find(&events)
	return events
}

func (r *eventRepository) FindByID(id string) (models.Event, error) {
	var event models.Event
	err := r.db.First(&event, "id = ?", id).Error

	return event, err
}

func (r *eventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Delete(event *models.Event) error {
	return r.db.Delete(event).Error
}
