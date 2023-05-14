package http

import (
	"easy-uni/config"
	"easy-uni/http/handlers"
	"easy-uni/http/middleware"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var (
	once sync.Once
	srv  *http.Server
)

func GetServer() *http.Server {
	once.Do(func() {

		log.Info().Msg("Initializing server")

		cfg := config.GetConfig()

		e := gin.Default()

		e.Use(middleware.CORS(cfg.AllowOrigin))

		r := e.Group("/api")

		routeUserHandler(r, cfg)

		s := &http.Server{
			Addr:    cfg.Port,
			Handler: e,
		}

		srv = s
	})

	return srv
}

func routeUserHandler(router *gin.RouterGroup, cfg config.Config) {
	// Users
	userHandler := handlers.GetUserHandler()
	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.Register)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/", userHandler.GetAll)
	userRouter.GET("/:user_id", userHandler.GetByID)

	userProtectedRouter := userRouter.Use(middleware.JwtAuth(cfg.ApiSecret))
	userProtectedRouter.GET("/current", userHandler.GetCurrent)

	// Universities
	universityHandler := handlers.GetUniversityHandler()
	universityRouter := router.Group("/universities")
	universityRouter.GET("/", universityHandler.Get)
	universityRouter.GET("/:university_id", universityHandler.GetByID)

	universityProtectedRouter := universityRouter.Use(middleware.JwtAuth(cfg.ApiSecret))
	universityProtectedRouter.POST("/", universityHandler.Create)
	universityProtectedRouter.PUT("/:university_id", universityHandler.Update)
	universityProtectedRouter.DELETE("/:university_id", universityHandler.Delete)

	// Faculties
	facultyHandler := handlers.GetFacultyHandler()
	facultyRouter := router.Group("/faculties")
	facultyRouter.GET("/", facultyHandler.Get)
	facultyRouter.GET("/:faculty_id", facultyHandler.GetByID)

	facultyProtectedRouter := facultyRouter.Use(middleware.JwtAuth(cfg.ApiSecret))
	facultyProtectedRouter.POST("/:university_id", facultyHandler.Create)
	facultyProtectedRouter.PUT("/:faculty_id", facultyHandler.Update)
	facultyProtectedRouter.DELETE("/:faculty_id", facultyHandler.Delete)

	// Events
	eventHandler := handlers.GetEventHandler()
	eventRouter := router.Group("/events")
	eventRouter.GET("/", eventHandler.GetAll)
	eventRouter.GET("/:event_id", eventHandler.GetByID)
	eventRouter.POST("/", eventHandler.Create)

	eventProtectedRouter := eventRouter.Use(middleware.JwtAuth(cfg.ApiSecret))
	eventProtectedRouter.DELETE("/:event_id", eventHandler.Delete)
}
