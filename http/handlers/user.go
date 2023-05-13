package handlers

import (
	"easy-uni/models"
	"easy-uni/services"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetCurrent(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	userService services.UserService
	mailService services.MailService
}

var (
	userOnce sync.Once
	userHnd  UserHandler
)

func GetUserHandler() UserHandler {
	userOnce.Do(func() {
		userHnd = &userHandler{
			userService: services.GetUserService(),
			mailService: services.GetMailService(),
		}
	})

	return userHnd
}

func (h *userHandler) GetAll(c *gin.Context) {
	users := h.userService.FindAll()
	c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetByID(c *gin.Context) {
	id := c.Param("user_id")
	user, err := h.userService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (h *userHandler) GetCurrent(c *gin.Context) {
	id := c.GetString("user_id")

	user, err := h.userService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Register(c *gin.Context) {

	var model models.RegisterUser
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Register(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mail := models.Mail{
		To:      []string{model.Email},
		Subject: "Welcome to EasyUni",
		Body:    fmt.Sprintf("Welcome %s", model.Name),
	}

	go h.mailService.Send(mail)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *userHandler) Login(c *gin.Context) {

	var model models.LoginUser
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
