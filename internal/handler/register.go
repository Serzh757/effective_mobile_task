package handler

import "github.com/gin-gonic/gin"

type MusicHandler struct {
}

func NewMusicHandler() *MusicHandler {
	return &MusicHandler{}
}

func RegisterRoutes(router *gin.Engine) {
	handler := NewMusicHandler()
	groupV1 := router.Group("/api/v1")
	groupV1.GET("/health", handler.HealthHandler)
}
