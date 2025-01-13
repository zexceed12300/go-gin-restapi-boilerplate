package errorhandler

import (
	"fmt"
	"go-gin-restapi-boilerplate/helpers"
	"go-gin-restapi-boilerplate/initializers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(c *gin.Context, err *error, types error) {
	var statusCode int

	traceId := time.Now().Unix()

	switch types.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	if err != nil {
		initializers.Logger.WithFields(logrus.Fields{
			"trace_id": traceId,
		}).Error((*err).Error())
	}

	res := helpers.Response(statusCode, helpers.ResponseParams{
		Message: fmt.Sprintf("%v: %v", traceId, types),
	})

	c.JSON(statusCode, res)
}
