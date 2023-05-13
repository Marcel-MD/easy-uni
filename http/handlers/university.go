package handlers

import (
	"easy-uni/models"
	"easy-uni/services"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type UniversityHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Get(c *gin.Context)

	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type universityHandler struct {
	service services.UniversityService
}

var (
	universityOnce sync.Once
	universityHnd  UniversityHandler
)

func GetUniversityHandler() UniversityHandler {
	universityOnce.Do(func() {
		universityHnd = &universityHandler{
			service: services.GetUniversityService(),
		}
	})

	return universityHnd
}

func (h *universityHandler) GetAll(c *gin.Context) {
	universities := h.service.FindAll()
	c.JSON(http.StatusOK, universities)
}

func (h *universityHandler) GetByID(c *gin.Context) {
	id := c.Param("university_id")
	university, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, university)
}

func (h *universityHandler) Get(c *gin.Context) {
	name := c.Query("name")
	country := c.Query("country")
	city := c.Query("city")

	universities := h.service.Find(name, country, city)
	c.JSON(http.StatusOK, universities)
}

func (h *universityHandler) Create(c *gin.Context) {
	var university models.CreateUniversity
	c.BindJSON(&university)

	newUniversity, err := h.service.Create(university)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUniversity)
}

func (h *universityHandler) Delete(c *gin.Context) {
	id := c.Param("university_id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "University deleted successfully"})
}
