package routes

import (
	"go-gin-restapi-boilerplate/handlers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(api *gin.RouterGroup) {
	api.GET("/image/:filename", handlers.FileGet)
}
