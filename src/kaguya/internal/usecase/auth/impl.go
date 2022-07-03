package auth

import (
	"context"
	jwtCommon "github.com/aryahmph/wattpad-clone/src/kaguya/common/jwt"
	mailProducer "github.com/aryahmph/wattpad-clone/src/kaguya/common/nsq/producer/mail"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/util"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
	userRepo "github.com/aryahmph/wattpad-clone/src/kaguya/internal/repository/user"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type authUsecaseImpl struct {
	userRepository userRepo.Repository
	jwtManager     *jwtCommon.JWTManager
	validate       *validator.Validate
	mailProducer   *mailProducer.MailProducer
}

func NewAuthUsecaseImpl(userRepository userRepo.Repository, jwtManager *jwtCommon.JWTManager,
	validate *validator.Validate, mailProducer *mailProducer.MailProducer) authUsecaseImpl {
	return authUsecaseImpl{userRepository: userRepository, jwtManager: jwtManager, validate: validate, mailProducer: mailProducer}
}

func (usecase authUsecaseImpl) Login(ctx context.Context, login userDomain.Login) (token string, err error) {
	err = usecase.validate.Struct(login)
	if err != nil {
		return token, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := usecase.userRepository.FindByUsername(ctx, login.Username)
	if err != nil {
		return token, err
	}

	if err = util.ComparePassword(user.PasswordHash, login.Password); err != nil {
		return token, err
	}

	token, err = usecase.jwtManager.GenerateToken(user.ID, time.Hour*24)
	if err != nil {
		return token, status.Error(codes.Internal, err.Error())
	}

	err = usecase.mailProducer.Publish(mailProducer.Receiver{
		To:      []string{user.Email},
		Subject: "Account security",
		Message: "There is login activity on your account!",
	})
	if err != nil {
		return token, err
	}

	return token, nil
}

func (usecase authUsecaseImpl) Session(ctx context.Context, token string) (user userDomain.User, err error) {
	if token == "" {
		return user, status.Error(codes.InvalidArgument, "token is required")
	}

	userID, err := usecase.jwtManager.VerifyToken(token)
	if err != nil {
		return user, err
	}

	return usecase.userRepository.FindByID(ctx, userID)
}
