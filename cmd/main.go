package main

import (
	"github.com/effective_mobile_task/infrastructure/router"
	"github.com/rs/zerolog/log"
)

func main() {
	api := router.InitRouter()

	log.Info().Msg("Server is running on port 8080")
	err := api.Run(":8080")
	if err != nil {
		panic(err)
	}
}
