package domain

import (
	"context"
)

type User struct {
	UserID      int64  `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedDate string `json:"created_date"`
}

type AuthUser struct {
	UserID      int64  `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// AuthUseCase represent the auth's usecases
type AuthUseCase interface {
	CreateUser(context.Context, *AuthUser) (User, error)
	GetUser(ctx context.Context) ([]User, error)
}

// UserRepository represent the auth's repository contract
type AuthRepository interface {
	CreateUser(ctx context.Context, user *AuthUser) (User, error)
	GetUser(ctx context.Context) ([]User, error)
}
