package repositories

import (
	"easy-uni/config"
	"easy-uni/models"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbOnce   sync.Once
	database *gorm.DB
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {

		log.Info().Msg("Initializing database")

		cfg := config.GetConfig()

		db, err := gorm.Open(postgres.Open(cfg.DatabaseUrl), &gorm.Config{})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
		}

		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.University{})
		db.AutoMigrate(&models.Faculty{})
		db.AutoMigrate(&models.Event{})

		database = db
	})

	return database
}

func CloseDB() error {
	dbSql, err := database.DB()
	if err != nil {
		return err
	}

	return dbSql.Close()
}
