package router

import (
	"github.com/gin-gonic/gin"

	"github.com/effective_mobile_task/config"
	"github.com/effective_mobile_task/internal/handler"
	"github.com/effective_mobile_task/internal/usecase"
)

func InitRouter(uc *usecase.SongUseCase, cfg *config.Config) *gin.Engine {
	router := gin.Default()
	handler.RegisterRoutes(router, uc, cfg)
	return router
}
