package errorhandler

import (
	"go-gin-restapi-boilerplate/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	res := helpers.Response(statusCode, helpers.ResponseParams{
		Message: err.Error(),
	})

	c.JSON(statusCode, res)
}
