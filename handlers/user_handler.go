package handlers

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserListHandler(c *gin.Context) {
	query := helpers.ListQueryParams{
		Limit: 20,
		Skip:  0,
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid query params",
		})
		return
	}

	users := []models.User{}

	db := initializers.DB.Limit(query.Limit)

	if query.Search != "" {
		db.Where("name LIKE ?", "%"+query.Search+"%")
	}
	if query.OrderBy != "" && query.OrderDir != "" {
		db.Order(query.OrderBy + " " + query.OrderDir)
	}
	if query.Skip != 0 {
		db.Offset(query.Skip)
	}

	db.Preload("Presences")

	db.Find(&users)

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived list of blud",
		Data:    users,
	})

	c.JSON(http.StatusOK, res)
}

func UserGetHandler(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	user := models.User{ID: uri.ID}

	db := initializers.DB.Model(&models.User{})

	db.Preload("Presences")

	if err := db.First(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.NotFoundError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived user",
		Data:    user,
	})

	c.JSON(http.StatusOK, res)
}

func UserCreateHandler(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	hashed, err := helpers.HashPassword(user.Password)

	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	user.Password = hashed

	if err := initializers.DB.Create(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusCreated, helpers.ResponseParams{
		Message: "Successfully created user",
		Data:    user,
	})

	c.JSON(http.StatusCreated, res)
}

func UserUpdateHandler(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	user := models.User{ID: uri.ID}

	if err := initializers.DB.First(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.NotFoundError{
			Message: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully updated user",
		Data:    user,
	})

	c.JSON(http.StatusOK, res)
}

func UserDeleteHandler(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	user := models.User{ID: uri.ID}

	if err := initializers.DB.First(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.NotFoundError{
			Message: err.Error(),
		})
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully deleted user",
	})

	c.JSON(http.StatusOK, res)
}
