package http

import (
	"encoding/json"
	"github.com/aryahmph/wattpad-clone/src/shirogane/common/grpc/pb"
	kaguyaGrpc "github.com/aryahmph/wattpad-clone/src/shirogane/internal/delivery/kaguya/grpc"
	kaguyaModel "github.com/aryahmph/wattpad-clone/src/shirogane/internal/model/kaguya"
	authUc "github.com/aryahmph/wattpad-clone/src/shirogane/internal/usecase/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const BEARER = len("Bearer ")

func MiddlewareAuth(service *kaguyaGrpc.KaguyaClient, authUsecase *authUc.AuthUsecaseImpl) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= BEARER {
			c.Error(status.Error(codes.InvalidArgument, "authorization header not valid"))
			c.Abort()
			return
		}

		tokenString := authHeader[BEARER:]

		ctx := c.Request.Context()
		session, err := authUsecase.GetSession(ctx, tokenString)
		if err != nil {
			res, err := service.AuthClient.Session(ctx, &pb.SessionRequest{
				Token: tokenString,
			})
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}

			response := kaguyaModel.CheckSessionResponse{
				ID:       res.GetId(),
				Username: res.GetUsername(),
				Email:    res.GetEmail(),
			}

			marshal, err := json.Marshal(response)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}

			session = string(marshal)
			err = authUsecase.SetSession(ctx, tokenString, string(marshal))
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}
		}

		c.Set("user_payload", session)
		c.Next()
	}
}
