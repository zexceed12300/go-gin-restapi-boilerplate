package main

import (
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"go-gin-restapi-boilerplate/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.NewDB()

	err := initializers.DB.AutoMigrate(&models.User{}, &models.Presence{}, &models.File{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

func main() {

	router := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	api := router.Group("/api")

	routes.AuthRoutes(api)
	routes.UserRoutes(api)
	routes.PresenceRoutes(api)
	routes.FileRoutes(api)

	router.Run()

}
