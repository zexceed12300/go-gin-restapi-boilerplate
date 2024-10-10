package routes

import (
	"go-gin-restapi-boilerplate/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup) {
	api.GET("/user", handlers.UserListHandler)
	api.GET("/user/:id", handlers.UserGetHandler)
	api.POST("/user", handlers.UserCreateHandler)
	api.PUT("/user/:id", handlers.UserUpdateHandler)
	api.DELETE("/user/:id", handlers.UserDeleteHandler)
}
