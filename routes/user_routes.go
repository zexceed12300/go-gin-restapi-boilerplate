package routes

import (
	"go-gin-restapi-boilerplate/handlers"
	"go-gin-restapi-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup) {
	api.GET("/user", middlewares.JWTMiddleware(), handlers.HandlerUserList)
	api.GET("/user/:id", middlewares.JWTMiddleware(), handlers.HandlerUserGet)
	api.POST("/user", middlewares.JWTMiddleware(), handlers.HandlerUserCreate)
	api.PUT("/user/:id", middlewares.JWTMiddleware(), handlers.HandlerUserUpdate)
	api.DELETE("/user/:id", middlewares.JWTMiddleware(), handlers.HandlerUserDelete)
}
