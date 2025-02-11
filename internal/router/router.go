package router

import (
	"Backend-trainee-assignment-winter-2025/internal/handler"
	"Backend-trainee-assignment-winter-2025/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler handler.Handler, secretKey []byte) *gin.Engine {
	router := gin.Default()

	router.POST("/api/auth", handler.NewUser)

	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware(secretKey))
	{
		authGroup.GET("/api/info", handler.Test)

	}

	return router
}
