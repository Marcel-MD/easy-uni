package handlers

import (
	"easy-uni/models"
	"easy-uni/services"
	"net/http"
	"sync"

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

func (h *eventHandler) GetAll(c *gin.Context) {
	events := h.service.FindAll()
	c.JSON(http.StatusOK, events)
}

func (h *eventHandler) GetByID(c *gin.Context) {
	id := c.Param("event_id")
	event, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

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

func (h *eventHandler) Delete(c *gin.Context) {
	id := c.Param("event_id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
