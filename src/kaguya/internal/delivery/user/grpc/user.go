package grpc

import (
	"context"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/grpc/pb"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/util"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
	userUc "github.com/aryahmph/wattpad-clone/src/kaguya/internal/usecase/user"
)

type GRPCUserServer struct {
	pb.UnimplementedGRPCUserServer
	userUsecase userUc.Usecase
}

func NewGRPCUserServer(userUsecase userUc.Usecase) *GRPCUserServer {
	return &GRPCUserServer{userUsecase: userUsecase}
}

func (delivery *GRPCUserServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	registerPayload := userDomain.Register{
		Username: request.GetUsername(),
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}
	user, err := delivery.userUsecase.Register(ctx, registerPayload)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{User: util.UserDomainToUserPb(user)}, nil
}
