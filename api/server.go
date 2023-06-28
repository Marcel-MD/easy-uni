package api

import (
	"net/http"
	"sync"

	"github.com/Marcel-MD/easy-uni/api/handlers"
	"github.com/Marcel-MD/easy-uni/api/middleware"
	"github.com/Marcel-MD/easy-uni/config"
	docs "github.com/Marcel-MD/easy-uni/docs"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		routeUniversityHandler(r, cfg)
		routeFacultyHandler(r, cfg)
		routeEventHandler(r, cfg)
		routeSwaggerHandler(r, cfg)

		s := &http.Server{
			Addr:    cfg.Port,
			Handler: e,
		}

		srv = s
	})

	return srv
}

func routeSwaggerHandler(router *gin.RouterGroup, cfg config.Config) {
	if cfg.Env == "prod" {
		return
	}

	docs.SwaggerInfo.Host = cfg.Host + cfg.Port
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func routeUserHandler(router *gin.RouterGroup, cfg config.Config) {
	// Users
	h := handlers.GetUserHandler()
	r := router.Group("/users")
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.GET("/", h.GetAll)
	r.GET("/:user_id", h.GetByID)

	pr := r.Use(middleware.JwtAuth(cfg.ApiSecret))
	pr.GET("/current", h.GetCurrent)
}

func routeUniversityHandler(router *gin.RouterGroup, cfg config.Config) {
	// Universities
	h := handlers.GetUniversityHandler()
	r := router.Group("/universities")
	r.GET("/", h.Get)
	r.GET("/:university_id", h.GetByID)

	pr := r.Use(middleware.JwtAuth(cfg.ApiSecret))
	pr.POST("/", h.Create)
	pr.PUT("/:university_id", h.Update)
	pr.DELETE("/:university_id", h.Delete)
}

func routeFacultyHandler(router *gin.RouterGroup, cfg config.Config) {
	// Faculties
	h := handlers.GetFacultyHandler()
	r := router.Group("/faculties")
	r.GET("/", h.Get)
	r.GET("/:faculty_id", h.GetByID)

	pr := r.Use(middleware.JwtAuth(cfg.ApiSecret))
	pr.POST("/:university_id", h.Create)
	pr.PUT("/:faculty_id", h.Update)
	pr.DELETE("/:faculty_id", h.Delete)
}

func routeEventHandler(router *gin.RouterGroup, cfg config.Config) {
	// Events
	h := handlers.GetEventHandler()
	r := router.Group("/events")
	r.GET("/", h.GetAll)
	r.GET("/:event_id", h.GetByID)
	r.POST("/", h.Create)

	pr := r.Use(middleware.JwtAuth(cfg.ApiSecret))
	pr.DELETE("/:event_id", h.Delete)
}
