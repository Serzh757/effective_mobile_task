package main

import (
	"github.com/rs/zerolog/log"

	"github.com/effective_mobile_task/config"
	"github.com/effective_mobile_task/infrastructure/db"
	"github.com/effective_mobile_task/infrastructure/router"
	"github.com/effective_mobile_task/internal/repository"
	"github.com/effective_mobile_task/internal/usecase"
)

func main() {
	cfg := config.NewConfig()
	database, err := db.InitPostgresDB(cfg)
	if err != nil {
		panic(err)
	}
	songRepository := repository.NewSongRepository(database)
	songUseCase := usecase.NewSongUseCase(songRepository)
	api := router.InitRouter(songUseCase)

	err = api.Run(":8080")
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Server is running on port 8080")
}
