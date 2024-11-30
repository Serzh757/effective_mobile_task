package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler /api/v1/health [GET]
// Проверка доступности сервиса
func (h *MusicHandler) HealthHandler(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(ctx.Writer, "status %d", http.StatusOK)
}
