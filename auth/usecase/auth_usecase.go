package usecase

import (
	"context"
	"log"
	"time"

	"github.com/questizen/core-system/domain"
	"github.com/questizen/core-system/helpers"
)

type AuthClaims struct {
	User *domain.User `json:"user"`
}

type authUseCase struct {
	authRepo       domain.AuthRepository
	contextTimeout time.Duration
}

// NewAuthUseCase will create new an authUseCase object representation of domain.AuthUsecase interface
func NewAuthUseCase(a domain.AuthRepository, timeout time.Duration) domain.AuthUseCase {
	return &authUseCase{
		authRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *authUseCase) CreateUser(c context.Context, user *domain.AuthUser) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	log.Println("run usecase create user", user)
	data, err := a.authRepo.CreateUser(ctx, user)

	helpers.PanicIfErr(err)

	return data, nil
}

func (a *authUseCase) GetUser(c context.Context) ([]domain.User, error) {
	data, err := a.authRepo.GetUser(c)
	helpers.PanicIfErr(err)

	log.Println("run usecase get users")
	return data, nil
}
