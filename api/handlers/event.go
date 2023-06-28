package handlers

import (
	"net/http"
	"sync"

	"github.com/Marcel-MD/easy-uni/models"
	"github.com/Marcel-MD/easy-uni/services"

	"github.com/gin-gonic/gin"
)

type EventHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type eventHandler struct {
	service services.EventService
}

var (
	eventOnce sync.Once
	eventHnd  EventHandler
)

func GetEventHandler() EventHandler {
	eventOnce.Do(func() {
		eventHnd = &eventHandler{
			service: services.GetEventService(),
		}
	})

	return eventHnd
}

// @Description get all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} models.Event
// @Router /events [get]
func (h *eventHandler) GetAll(c *gin.Context) {
	events := h.service.FindAll()
	c.JSON(http.StatusOK, events)
}

// @Description get event by id
// @Tags events
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} models.Event
// @Router /events/{event_id} [get]
func (h *eventHandler) GetByID(c *gin.Context) {
	id := c.Param("event_id")
	event, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// @Description create event
// @Tags events
// @Accept json
// @Produce json
// @Param event body models.CreateEvent true "Event"
// @Success 201 {object} models.Event
// @Router /events [post]
func (h *eventHandler) Create(c *gin.Context) {
	var event models.CreateEvent
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent, err := h.service.Create(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

// @Description delete event
// @Tags events
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Security ApiKeyAuth
// @Success 204
// @Router /events/{event_id} [delete]
func (h *eventHandler) Delete(c *gin.Context) {
	id := c.Param("event_id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
