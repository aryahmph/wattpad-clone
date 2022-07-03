package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type HTTPServer struct {
	Router *gin.Engine
}

func init() {
	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}

	gin.DisableConsoleColor()
}

func NewHTTPServer() HTTPServer {
	router := gin.Default()

	return HTTPServer{
		Router: router,
	}
}
