package handlers

import (
	"easy-uni/models"
	"easy-uni/services"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type FacultyHandler interface {
	GetAll(c *gin.Context)
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

func (h *facultyHandler) GetAll(c *gin.Context) {
	faculties := h.service.FindAll()
	c.JSON(http.StatusOK, faculties)
}

func (h *facultyHandler) GetByID(c *gin.Context) {
	id := c.Param("faculty_id")
	faculty, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, faculty)
}

func (h *facultyHandler) Get(c *gin.Context) {
	name := c.Query("name")
	country := c.Query("country")
	city := c.Query("city")
	domain := c.Query("domain")
	budgetStr := c.Query("budget")
	budget := -1

	if budgetStr != "" {
		var err error
		budget, err = strconv.Atoi(budgetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	faculties := h.service.Find(name, country, city, domain, budget)
	c.JSON(http.StatusOK, faculties)
}

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

func (h *facultyHandler) Delete(c *gin.Context) {
	id := c.Param("faculty_id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Faculty deleted successfully"})
}
