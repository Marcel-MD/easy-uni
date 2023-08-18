package repositories

import (
	"sync"

	"github.com/Marcel-MD/easy-uni/data"
	"github.com/Marcel-MD/easy-uni/models"

	"github.com/rs/zerolog/log"
)

type EventRepository interface {
	FindAll(page, size int) ([]models.Event, error)
	FindById(id string) (models.Event, error)
	Create(t *models.Event) error
	Update(t *models.Event) error
	Delete(t *models.Event) error
}

type eventRepository struct {
	data.Repository[models.Event]
}

var (
	eventOnce sync.Once
	eventRepo EventRepository
)

func GetEventRepository() EventRepository {
	eventOnce.Do(func() {
		log.Info().Msg("Initializing event repository")
		eventRepo = &eventRepository{
			Repository: data.NewRepository[models.Event](),
		}
	})
	return eventRepo
}
