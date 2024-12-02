package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/effective_mobile_task/config"
	errors "github.com/effective_mobile_task/internal/middleware"
	"github.com/effective_mobile_task/internal/swagger"
	"github.com/effective_mobile_task/internal/usecase"
	"github.com/effective_mobile_task/internal/view"
)

type MusicHandler struct {
	cfg *config.Config
	uc  *usecase.SongUseCase
}

func NewMusicHandler(uc *usecase.SongUseCase, cfg *config.Config) *MusicHandler {
	return &MusicHandler{
		uc:  uc,
		cfg: cfg,
	}
}

func RegisterRoutes(router *gin.Engine, uc *usecase.SongUseCase, cfg *config.Config) {
	router.Use(errors.Wrapper)
	groupV1 := router.Group("/api/v1")
	handler := NewMusicHandler(uc, cfg)
	_ = initializeSwagger(router)

	groupV1.GET("/health", handler.HealthHandler)

	// Получение всех песен
	groupV1.GET("/songs", handler.AllSongs)
	// Получение песни по ИД
	groupV1.GET("/song/:id", handler.SongByID)
	// Добавление новой песни
	groupV1.POST("/song", handler.AddSong)
	// Изменение текущей песни
	groupV1.PUT("/song", handler.UpdateSong)
	// Удаление песни
	groupV1.DELETE("/song/:id", handler.DeleteSongByID)
}

// initializeSwagger - инициализация документации
func initializeSwagger(router *gin.Engine) error {
	err := swagger.Register(router.Group("/api/v1"), view.GetSwagger, "/api/v1")
	if err != nil {
		return err
	}
	return nil
}
