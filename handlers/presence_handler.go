package handlers

import (
	"go-gin-restapi-boilerplate/errorhandler"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"go-gin-restapi-boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PresenceListHandler(c *gin.Context) {
	userID := c.GetInt("userID")

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

	presences := []models.Presence{}

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

	db.Joins("User")

	if err := db.Find(&presences, &models.Presence{UserID: userID}).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived list of presence",
		Data:    presences,
	})

	c.JSON(http.StatusOK, res)
}

func PresenceGetHandler(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	presence := models.Presence{ID: uri.ID}

	db := initializers.DB.Model(&models.Presence{})

	db.Joins("User")

	if err := db.First(&presence).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully retrived presence",
		Data:    presence,
	})

	c.JSON(http.StatusOK, res)
}

func PresenceCreateHandler(c *gin.Context) {
	req := models.PresenceRequest{}

	if err := c.ShouldBind(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	presence := models.Presence{
		UserID:   req.UserID,
		Location: req.Location,
		Type:     req.Type,
	}

	if req.Image != nil {
		filename, err := UploadFile(req.Image)
		if err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
				Message: err.Error(),
			})
			return
		}

		presence.Image = &filename
	}

	if err := initializers.DB.Create(&presence).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully created presence",
		Data:    presence,
	})

	c.JSON(http.StatusOK, res)
}

func PresenceUpdateHandler(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	presence := models.Presence{}

	if err := initializers.DB.First(&presence, &models.Presence{ID: uri.ID}).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	req := models.PresenceRequest{
		UserID:   presence.UserID,
		Location: presence.Location,
		Type:     presence.Type,
	}

	if err := c.ShouldBind(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	update := models.Presence{
		UserID:   req.UserID,
		Location: req.Location,
		Type:     req.Type,
	}

	if req.Image != nil {
		filename, err := UploadFile(req.Image)
		if err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
				Message: err.Error(),
			})
			return
		}

		update.Image = &filename
	}

	if err := initializers.DB.Model(&presence).Updates(&update).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully updated presence",
		Data:    presence,
	})

	c.JSON(http.StatusOK, res)
}

func PresenceDeleteHandler(c *gin.Context) {
	uri := helpers.UriIDParam{}

	if err := c.ShouldBindUri(&uri); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "Invalid id",
		})
		return
	}

	presence := models.Presence{ID: uri.ID}

	if err := initializers.DB.First(&presence).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	if err := initializers.DB.Delete(&presence).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully deleted presence",
		Data:    presence,
	})

	c.JSON(http.StatusOK, res)
}

func PresenceUserHandler(c *gin.Context) {
	userID := c.GetInt("userID")

	req := models.PresenceRequest{UserID: userID}

	if err := c.ShouldBind(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	presence := models.Presence{
		UserID:   userID,
		Location: req.Location,
		Type:     req.Type,
	}

	if req.Image != nil {
		filename, err := UploadFile(req.Image)
		if err != nil {
			errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
				Message: err.Error(),
			})
			return
		}

		presence.Image = &filename
	}

	if err := initializers.DB.Create(&presence).Error; err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: err.Error(),
		})
		return
	}

	res := helpers.Response(http.StatusOK, helpers.ResponseParams{
		Message: "Successfully created presence",
		Data:    presence,
	})

	c.JSON(http.StatusOK, res)
}
