package routes

import (
	"go-gin-restapi-boilerplate/handlers"
	"go-gin-restapi-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func PresenceRoutes(api *gin.RouterGroup) {
	api.GET("/presence", handlers.PresenceListHandler)
	api.GET("/presence/:id", handlers.PresenceGetHandler)
	api.POST("/presence", handlers.PresenceCreateHandler)
	api.PUT("/presence/:id", handlers.PresenceUpdateHandler)
	api.DELETE("/presence/:id", handlers.PresenceDeleteHandler)

	api.Use(middlewares.JWTMiddleware())
	{
		api.POST("/user/presence", handlers.PresenceUserHandler)
	}
}
