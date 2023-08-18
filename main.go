package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Marcel-MD/easy-uni/api"
	"github.com/Marcel-MD/easy-uni/data"

	"github.com/rs/zerolog/log"
)

// @title Easy-Uni API
// @description This is the API for the Easy-Uni application.
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "bearer" followed by a space and JWT token
func main() {
	srv := api.GetServer()

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
		log.Info().Msg("All server connections are closed")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	<-quit
	log.Warn().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	if err := data.CloseDB(); err != nil {
		log.Fatal().Err(err).Msg("Failed to close db connection")
	}

	log.Info().Msg("Server exited properly")
}
