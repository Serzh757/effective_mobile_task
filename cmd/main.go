package main

import (
	"github.com/rs/zerolog/log"

	"github.com/effective_mobile_task/infrastructure/router"
)

func main() {
	api := router.InitRouter()

	log.Info().Msg("Server is running on port 8080")
	err := api.Run(":8080")
	if err != nil {
		panic(err)
	}
}
