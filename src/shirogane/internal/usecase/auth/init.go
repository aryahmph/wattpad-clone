package auth

import (
	"context"
	"fmt"
	authRepo "github.com/aryahmph/wattpad-clone/src/shirogane/internal/repository/auth"
	"time"
)

type AuthUsecaseImpl struct {
	authRepository authRepo.Repository
}

func NewAuthUsecaseImpl(authRepository authRepo.Repository) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{authRepository: authRepository}
}

func (usecase *AuthUsecaseImpl) SetSession(ctx context.Context, token, data string) (err error) {
	return usecase.authRepository.Set(ctx, fmt.Sprintf("auth_%s", token), data, 1*time.Minute)
}

func (usecase *AuthUsecaseImpl) GetSession(ctx context.Context, token string) (data string, err error) {
	data, err = usecase.authRepository.Get(ctx, fmt.Sprintf("auth_%s", token))
	if err != nil {
		return "", err
	}
	return data, nil
}
