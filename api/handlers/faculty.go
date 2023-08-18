package handlers

import (
	"net/http"
	"sync"

	"github.com/Marcel-MD/easy-uni/models"
	"github.com/Marcel-MD/easy-uni/services"

	"github.com/gin-gonic/gin"
)

type FacultyHandler interface {
	GetByID(c *gin.Context)
	Get(c *gin.Context)

	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type facultyHandler struct {
	service services.FacultyService
}

var (
	facultyOnce sync.Once
	facultyHnd  FacultyHandler
)

func GetFacultyHandler() FacultyHandler {
	facultyOnce.Do(func() {
		facultyHnd = &facultyHandler{
			service: services.GetFacultyService(),
		}
	})

	return facultyHnd
}

// @Description get faculty by id
// @Tags faculties
// @Accept json
// @Produce json
// @Param faculty_id path string true "Faculty ID"
// @Success 200 {object} models.Faculty
// @Router /faculties/{faculty_id} [get]
func (h *facultyHandler) GetByID(c *gin.Context) {
	id := c.Param("faculty_id")
	faculty, err := h.service.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, faculty)
}

// @Description get faculty by name, country, city, domain, budget
// @Tags faculties
// @Accept json
// @Produce json
// @Param faculty query models.FacultyQuery false "Faculty"
// @Success 200 {array} models.Faculty
// @Router /faculties [get]
func (h *facultyHandler) Get(c *gin.Context) {
	var query models.FacultyQuery
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	faculties := h.service.Find(query)
	c.JSON(http.StatusOK, faculties)
}

// @Description create faculty
// @Tags faculties
// @Accept json
// @Produce json
// @Param university_id path string true "University ID"
// @Param user body models.CreateFaculty true "Faculty"
// @Security ApiKeyAuth
// @Success 200 {object} models.Faculty
// @Router /faculties/{university_id} [post]
func (h *facultyHandler) Create(c *gin.Context) {
	universityID := c.Param("university_id")
	var faculty models.CreateFaculty
	err := c.BindJSON(&faculty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newFaculty, err := h.service.Create(universityID, faculty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newFaculty)
}

// @Description update faculty
// @Tags faculties
// @Accept json
// @Produce json
// @Param faculty_id path string true "Faculty ID"
// @Param user body models.CreateFaculty true "Faculty"
// @Security ApiKeyAuth
// @Success 200 {object} models.Faculty
// @Router /faculties/{faculty_id} [put]
func (h *facultyHandler) Update(c *gin.Context) {
	id := c.Param("faculty_id")
	var faculty models.CreateFaculty
	err := c.BindJSON(&faculty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFaculty, err := h.service.Update(id, faculty)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedFaculty)
}

// @Description delete faculty
// @Tags faculties
// @Accept json
// @Produce json
// @Param faculty_id path string true "Faculty ID"
// @Security ApiKeyAuth
// @Success 200 {string} string "Faculty deleted successfully"
// @Router /faculties/{faculty_id} [delete]
func (h *facultyHandler) Delete(c *gin.Context) {
	id := c.Param("faculty_id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Faculty deleted successfully"})
}
