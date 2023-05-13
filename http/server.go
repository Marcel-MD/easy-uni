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
	uh := handlers.GetUserHandler()

	ug := router.Group("/users")
	ug.POST("/register", uh.Register)
	ug.POST("/login", uh.Login)
	ug.GET("/", uh.GetAll)
	ug.GET("/:user_id", uh.GetById)

	ur := ug.Use(middleware.JwtAuth(cfg.ApiSecret))
	ur.GET("/current", uh.GetCurrent)
}
