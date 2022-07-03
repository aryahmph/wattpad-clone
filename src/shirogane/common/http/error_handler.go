package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func MiddlewareErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs[0]
			if !err.IsType(gin.ErrorTypePrivate) {
				var ves validator.ValidationErrors
				if errors.As(err, &ves) {
					keys := make(map[string]string)
					for _, ve := range ves {
						keys[ve.Field()] = ve.Tag()
					}
					c.JSON(c.Writer.Status(), Error{
						Code:    c.Writer.Status(),
						Message: err.Error(),
						Errors:  keys,
					})
					return
				}
			}

			fmt.Println(err)
			if e, ok := status.FromError(err.Err); ok {
				switch e.Code() {
				case codes.InvalidArgument:
					c.JSON(http.StatusBadRequest, Error{
						Code:    http.StatusBadRequest,
						Message: e.Message(),
					})
				case codes.AlreadyExists:
					c.JSON(http.StatusBadRequest, Error{
						Code:    http.StatusBadRequest,
						Message: e.Message(),
					})
				case codes.NotFound:
					c.JSON(http.StatusNotFound, Error{
						Code:    http.StatusNotFound,
						Message: e.Message(),
					})
				case codes.Unauthenticated:
					c.JSON(http.StatusUnauthorized, Error{
						Code:    http.StatusUnauthorized,
						Message: e.Message(),
					})
				default:
					c.JSON(http.StatusInternalServerError, Error{
						Code:    http.StatusInternalServerError,
						Message: "Internal server error",
					})
				}
			} else {
				c.JSON(http.StatusInternalServerError, Error{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				})
			}
		}
	}
}
