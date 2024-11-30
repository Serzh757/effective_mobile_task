package router

import (
	"github.com/effective_mobile_task/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	handler.RegisterRoutes(router)
	return router
}
