package http

import (
	httpCommon "github.com/aryahmph/wattpad-clone/src/shirogane/common/http"
	kaguyaGrpc "github.com/aryahmph/wattpad-clone/src/shirogane/internal/delivery/kaguya/grpc"
	authUc "github.com/aryahmph/wattpad-clone/src/shirogane/internal/usecase/auth"
	"github.com/gin-gonic/gin"
)

func RegisterKaguyaService(router *gin.RouterGroup, addr string, authUsecase *authUc.AuthUsecaseImpl) (*kaguyaGrpc.KaguyaClient, error) {
	kaguyaClient, err := kaguyaGrpc.NewKaguyaClient(addr)
	if err != nil {
		return nil, err
	}
	user := router.Group("/user")
	user.POST("", kaguyaClient.Register)

	user.Use(httpCommon.MiddlewareAuth(kaguyaClient, authUsecase))
	user.GET("/ping", kaguyaClient.Ping)

	auth := router.Group("/auth")
	auth.POST("", kaguyaClient.Login)

	return kaguyaClient, nil
}
