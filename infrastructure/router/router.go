package router

import (
	"github.com/gin-gonic/gin"

	"github.com/effective_mobile_task/internal/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	handler.RegisterRoutes(router)
	return router
}
