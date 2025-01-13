package routes

import (
	"go-gin-restapi-boilerplate/handlers"
	"go-gin-restapi-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(api *gin.RouterGroup) {
	router := api.Group("/auth")

	router.POST("/login", handlers.HandlerAuthLogin)
	router.GET("/user", middlewares.JWTMiddleware(), handlers.HandlerAuthUserMe)
}
