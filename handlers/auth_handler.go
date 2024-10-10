package handlers

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthUserMeHandler(c *gin.Context) {
	userID := c.GetInt("userID")

	var user models.User

	db := initializers.DB.Model(&models.User{})

	db.Preload("Presences")

	db.Find(&user, userID)

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived user",
		Data:    user,
	})

	c.JSON(http.StatusOK, res)
}

func AuthLoginHandler(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	var user models.User
	initializers.DB.Where("username = ?", req.Username).First(&user)

	if !helpers.CheckPasswordHash(req.Password, user.Password) {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "Email or password invalid",
		})
		return
	}

	access_token, err := helpers.GenerateAccessToken(&user)
	if err != nil {
		log.Printf("could not generate access_token %s", err.Error())
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: err.Error(),
		})
		return
	}

	result := models.LoginResponse{
		AccessToken: access_token,
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully logged in",
		Data:    result,
	})

	c.JSON(http.StatusOK, res)
}
