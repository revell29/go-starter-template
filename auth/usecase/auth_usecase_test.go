package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	ucase "github.com/questizen/core-system/auth/usecase"
	"github.com/questizen/core-system/domain"
	mocks "github.com/questizen/core-system/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAuthFlow(t *testing.T) {
	mockAuthRepo := new(mocks.AuthRepository)
	mockAuth := domain.AuthUser{
		Name:        "Apsya",
		Email:       "diraapsya9@gmail.com",
		PhoneNumber: "123456789",
	}

	t.Run("Create User", func(t *testing.T) {
		tempMockAuth := mockAuth
		tempMockUser := domain.User{
			UserID:      1,
			Name:        "Apsya",
			Email:       "diraapsya9@gmail.com",
			PhoneNumber: "123456789",
		}
		mockAuthRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.AuthUser")).Return(tempMockUser, nil).Once()

		u := ucase.NewAuthUseCase(mockAuthRepo, time.Second*5)
		_, err := u.CreateUser(context.TODO(), &tempMockAuth)

		assert.NoError(t, err)
		assert.Equal(t, mockAuth.Email, tempMockAuth.Email)
		mockAuthRepo.AssertExpectations(t)
	})
}

func TestGetUser(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)

	require.NoError(t, err)
	mockAuthRepo := new(mocks.AuthRepository)
	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)
	mockAuthRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListUser, "", nil).Once()
}
