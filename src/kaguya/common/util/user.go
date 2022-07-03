package util

import (
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/grpc/pb"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
)

func UserDomainToUserPb(user userDomain.User) *pb.User {
	return &pb.User{
		Id:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}
