package handlers

import (
	"net/http"
	"sync"

	"github.com/Marcel-MD/easy-uni/models"
	"github.com/Marcel-MD/easy-uni/services"

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
}

var (
	userOnce sync.Once
	userHnd  UserHandler
)

func GetUserHandler() UserHandler {
	userOnce.Do(func() {
		userHnd = &userHandler{
			userService: services.GetUserService(),
		}
	})

	return userHnd
}

// @Description get all users
// @Tags users
// @Accept json
// @Produce json
// @Param pagination query models.PaginationQuery false "Pagination"
// @Success 200 {array} models.User
// @Router /users [get]
func (h *userHandler) GetAll(c *gin.Context) {
	query := models.PaginationQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.userService.FindAll(query.Page, query.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Description get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{user_id} [get]
func (h *userHandler) GetByID(c *gin.Context) {
	id := c.Param("user_id")
	user, err := h.userService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// @Description get current user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Router /users/current [get]
func (h *userHandler) GetCurrent(c *gin.Context) {
	id := c.GetString("user_id")

	user, err := h.userService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Description register user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.RegisterUser true "User"
// @Success 200 {object} models.Token
// @Router /users/register [post]
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

	c.JSON(http.StatusOK, models.Token{Token: token})
}

// @Description login user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.LoginUser true "User"
// @Success 200 {object} models.Token
// @Router /users/login [post]
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

	c.JSON(http.StatusOK, models.Token{Token: token})
}
