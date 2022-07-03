package user

import (
	"context"
	mailProducer "github.com/aryahmph/wattpad-clone/src/kaguya/common/nsq/producer/mail"
	"github.com/aryahmph/wattpad-clone/src/kaguya/common/util"
	userDomain "github.com/aryahmph/wattpad-clone/src/kaguya/internal/domain/user"
	userRepo "github.com/aryahmph/wattpad-clone/src/kaguya/internal/repository/user"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userUsecaseImpl struct {
	userRepository userRepo.Repository
	validate       *validator.Validate
	mailProducer   *mailProducer.MailProducer
}

func NewUserUsecaseImpl(userRepository userRepo.Repository, validate *validator.Validate,
	mailProducer *mailProducer.MailProducer) userUsecaseImpl {
	return userUsecaseImpl{userRepository: userRepository, validate: validate, mailProducer: mailProducer}
}

func (usecase userUsecaseImpl) Register(ctx context.Context, registrar userDomain.Register) (userDomain.User, error) {
	err := usecase.validate.Struct(registrar)
	if err != nil {
		return userDomain.User{}, status.Error(codes.InvalidArgument, err.Error())
	}

	_, err = usecase.userRepository.FindByUsername(ctx, registrar.Username)
	if err == nil {
		return userDomain.User{}, status.Error(codes.AlreadyExists, "username already exists")
	}

	_, err = usecase.userRepository.FindByEmail(ctx, registrar.Email)
	if err == nil {
		return userDomain.User{}, status.Error(codes.AlreadyExists, "email already exists")
	}

	hashPassword, err := util.HashPassword(registrar.Password)
	if err != nil {
		return userDomain.User{}, status.Error(codes.Internal, err.Error())
	}

	user := userDomain.User{
		Username:     registrar.Username,
		Email:        registrar.Email,
		PasswordHash: hashPassword,
	}

	rid, err := usecase.userRepository.Insert(ctx, user)
	if err != nil {
		return userDomain.User{}, err
	}
	user.ID = rid

	err = usecase.mailProducer.Publish(mailProducer.Receiver{
		To:      []string{registrar.Email},
		Subject: "Account registration successful",
		Message: "Thank you for registering on our platform!",
	})
	if err != nil {
		return userDomain.User{}, err
	}

	return user, nil
}
