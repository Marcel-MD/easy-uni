package handlers

import (
	"easy-uni/models"
	"easy-uni/services"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type UniversityHandler interface {
	GetByID(c *gin.Context)
	Get(c *gin.Context)

	Create(c *gin.Context)
	Update(c *gin.Context)
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

// @Description get university by id
// @Tags universities
// @Accept json
// @Produce json
// @Param university_id path string true "University ID"
// @Success 200 {object} models.University
// @Router /universities/{university_id} [get]
func (h *universityHandler) GetByID(c *gin.Context) {
	id := c.Param("university_id")
	university, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, university)
}

// @Description get university by name, country, city
// @Tags universities
// @Accept json
// @Produce json
// @Param name query string false "University Name"
// @Param country query string false "University Country"
// @Param city query string false "University City"
// @Success 200 {array} models.University
// @Router /universities [get]
func (h *universityHandler) Get(c *gin.Context) {
	name := c.Query("name")
	country := c.Query("country")
	city := c.Query("city")

	universities := h.service.Find(name, country, city)
	c.JSON(http.StatusOK, universities)
}

// @Description create university
// @Tags universities
// @Accept json
// @Produce json
// @Param university body models.CreateUniversity true "University"
// @Security ApiKeyAuth
// @Success 200 {object} models.University
// @Router /universities [post]
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

// @Description update university
// @Tags universities
// @Accept json
// @Produce json
// @Param university_id path string true "University ID"
// @Param university body models.CreateUniversity true "University"
// @Security ApiKeyAuth
// @Success 200 {object} models.University
// @Router /universities/{university_id} [put]
func (h *universityHandler) Update(c *gin.Context) {
	id := c.Param("university_id")
	var university models.CreateUniversity
	c.BindJSON(&university)

	updatedUniversity, err := h.service.Update(id, university)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUniversity)
}

// @Description delete university
// @Tags universities
// @Accept json
// @Produce json
// @Param university_id path string true "University ID"
// @Security ApiKeyAuth
// @Success 200 {string} string "University deleted successfully"
// @Router /universities/{university_id} [delete]
func (h *universityHandler) Delete(c *gin.Context) {
	id := c.Param("university_id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "University deleted successfully"})
}
