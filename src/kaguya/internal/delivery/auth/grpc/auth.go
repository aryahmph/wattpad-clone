package grpc

import (
	"context"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/grpc/pb"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
	authUc "github.com/aryahmph/wattpad-clone/src/kaguya/internal/usecase/auth"
)

type GRPCAuthServer struct {
	pb.UnimplementedGRPCAuthServer
	authUsecase authUc.Usecase
}

func NewGRPCAuthServer(authUsecase authUc.Usecase) *GRPCAuthServer {
	return &GRPCAuthServer{authUsecase: authUsecase}
}

func (delivery *GRPCAuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginPayload := userDomain.Login{
		Username: request.GetUsername(),
		Password: request.GetPassword(),
	}

	token, err := delivery.authUsecase.Login(ctx, loginPayload)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}

func (delivery *GRPCAuthServer) Session(ctx context.Context, request *pb.SessionRequest) (*pb.SessionResponse, error) {
	user, err := delivery.authUsecase.Session(ctx, request.GetToken())
	if err != nil {
		return nil, err
	}
	return &pb.SessionResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
