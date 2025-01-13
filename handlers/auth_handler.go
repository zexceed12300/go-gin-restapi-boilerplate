package handlers

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"go-gin-restapi-boilerplate/models/validation"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerAuthUserMe(c *gin.Context) {
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

func HandlerAuthLogin(c *gin.Context) {
	var req validation.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	var user *models.User

	if err := initializers.DB.First(&user, models.User{
		Email: req.Email,
	}); err.Error != nil {
		errorhandler.ErrorHandler(c, &err.Error, &errorhandler.UnauthorizedError{
			Message: "Email or password invalid",
		})
		return
	}

	if !helpers.CheckPasswordHash(req.Password, user.PasswordHash) {
		errorhandler.ErrorHandler(c, nil, &errorhandler.UnauthorizedError{
			Message: "Email or password invalid",
		})
		return
	}

	access_token, err := helpers.GenerateAccessToken(user)
	if err != nil {
		log.Printf("could not generate access_token %s", err.Error())
		errorhandler.ErrorHandler(c, &err, &errorhandler.UnauthorizedError{
			Message: "Something's error",
		})
		return
	}

	refresh_token, err := helpers.GenerateRefreshToken(user)
	if err != nil {
		log.Printf("could not generate access_token %s", err.Error())
		errorhandler.ErrorHandler(c, &err, &errorhandler.UnauthorizedError{
			Message: "Something's error",
		})
		return
	}

	user.AccessToken = access_token
	user.RefreshToken = refresh_token

	if err := initializers.DB.Save(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Something's error",
		})
		return
	}

	result := validation.LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully login",
		Data:    result,
	})

	c.JSON(http.StatusOK, res)
}
