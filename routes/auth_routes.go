package routes

import (
	"go-gin-restapi-boilerplate/handlers"
	"go-gin-restapi-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	router := api.Group("/auth")

	router.POST("/login", handlers.AuthLoginHandler)

	router.Use(middlewares.JWTMiddleware())
	{
		router.GET("/me", handlers.AuthUserMeHandler)
	}
}
