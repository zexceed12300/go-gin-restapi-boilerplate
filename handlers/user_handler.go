package handlers

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerUserList(c *gin.Context) {
	query := helpers.ListQueryParams{
		Limit: 20,
		Skip:  0,
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Invalid query params",
		})
		return
	}

	users := []models.User{}
	var recordsTotal int64

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

	if err := db.Find(&users).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Failed to retrive list of users",
		})
		return
	}

	if err := db.Offset(-1).Limit(-1).Count(&recordsTotal).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Failed to retrive total records",
		})
		return
	}

	data := struct {
		Users        []models.User `json:"users"`
		RecordsTotal int64         `json:"recordsTotal"`
	}{
		Users:        users,
		RecordsTotal: recordsTotal,
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived list of blud",
		Data:    data,
	})

	c.JSON(http.StatusOK, res)
}

func HandlerUserGet(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	user := models.User{ID: uri.ID}

	db := initializers.DB.Model(&models.User{})

	db.Preload("Presences")

	if err := db.First(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.NotFoundError{
			Message: "User not found",
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived user",
		Data:    user,
	})

	c.JSON(http.StatusOK, res)
}

func HandlerUserCreate(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Invalid request body",
		})
		return
	}

	hashed, err := helpers.HashPassword(user.PasswordHash)

	if err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Failed to hash password",
		})
		return
	}

	user.PasswordHash = hashed

	if err := initializers.DB.Create(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Failed to create user",
		})
		return
	}

	res := helpers.Response(http.StatusCreated, helpers.ResponseParams{
		Message: "Successfully created user",
		Data:    user,
	})

	c.JSON(http.StatusCreated, res)
}

func HandlerUserUpdate(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	user := models.User{ID: uri.ID}

	if err := initializers.DB.First(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.NotFoundError{
			Message: "User not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Invalid request body",
		})
		return
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Failed to update user",
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully updated user",
		Data:    user,
	})

	c.JSON(http.StatusOK, res)
}

func HandlerUserDelete(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	user := models.User{ID: uri.ID}

	if err := initializers.DB.First(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.NotFoundError{
			Message: "User not found",
		})
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		errorhandler.ErrorHandler(c, &err, &errorhandler.InternalServerError{
			Message: "Failed to delete user",
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully deleted user",
	})

	c.JSON(http.StatusOK, res)
}
