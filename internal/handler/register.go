package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/effective_mobile_task/internal/swagger"
	"github.com/effective_mobile_task/internal/view"
)

type MusicHandler struct{}

func NewMusicHandler() *MusicHandler {
	return &MusicHandler{}
}

func RegisterRoutes(router *gin.Engine) {
	groupV1 := router.Group("/api/v1")
	handler := NewMusicHandler()
	_ = initializeSwagger(router)

	groupV1.GET("/health", handler.HealthHandler)

	// Получение всех песен
	groupV1.GET("/songs", handler.HealthHandler)
	// Получение песни по ИД
	groupV1.GET("/song/:id", handler.HealthHandler)
	// Добавление новой песни
	groupV1.POST("/song", handler.HealthHandler)
	// Изменение текущей песни
	groupV1.PUT("/song", handler.HealthHandler)
	// Удаление песни
	groupV1.DELETE("/song/:id", handler.HealthHandler)
}

// initializeSwagger - инициализация документации
func initializeSwagger(router *gin.Engine) error {
	err := swagger.Register(router.Group("/api/v1"), view.GetSwagger, "/api/v1")
	if err != nil {
		return err
	}
	return nil
}
