package grpc

import (
	"github.com/aryahmph/wattpad-clone/src/shirogane/common/grpc/pb"
	kaguyaModel "github.com/aryahmph/wattpad-clone/src/shirogane/internal/model/kaguya"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type KaguyaClient struct {
	UserClient pb.GRPCUserClient
	AuthClient pb.GRPCAuthClient
}

func NewKaguyaClient(addr string) (*KaguyaClient, error) {
	clientConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userClient := pb.NewGRPCUserClient(clientConn)
	authClient := pb.NewGRPCAuthClient(clientConn)
	return &KaguyaClient{UserClient: userClient, AuthClient: authClient}, nil
}

func (svc *KaguyaClient) Register(c *gin.Context) {
	var requestBody kaguyaModel.RegisterRequest
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	ctx := c.Request.Context()
	response, err := svc.UserClient.Register(ctx, &pb.RegisterRequest{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id": response.User.GetId(),
		},
	})
}

func (svc *KaguyaClient) Login(c *gin.Context) {
	var requestBody kaguyaModel.LoginRequest
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	ctx := c.Request.Context()
	response, err := svc.AuthClient.Login(ctx, &pb.LoginRequest{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": response.GetToken(),
		},
	})
}

func (svc *KaguyaClient) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": c.MustGet("user_payload"),
		},
	})
}
