package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	authHttp "github.com/questizen/core-system/auth/delivery/http"
	"github.com/questizen/core-system/domain"
	"github.com/questizen/core-system/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	mockUser := &domain.AuthUser{
		Name:        "Apsya",
		Email:       "diraapsya9@gmail.com",
		PhoneNumber: "123456789",
	}

	tempMockUser := mockUser
	tempMockResult := domain.User{
		UserID:      1,
		Name:        "Apsya",
		Email:       "diraapsya@gmail.com",
		PhoneNumber: "123456789",
	}
	mockUseCase := new(mocks.AuthUseCase)

	body, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUseCase.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.AuthUser")).Return(tempMockResult, nil).Once()

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/signup", strings.NewReader(string(body)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signup")

	handler := authHttp.Handler{
		AUsecase: mockUseCase,
	}
	err = handler.CreateUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)

	require.NoError(t, err)
	mockUseCase := new(mocks.AuthUseCase)
	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	mockUseCase.On("GetUser", mock.Anything).Return(mockListUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := authHttp.Handler{
		AUsecase: mockUseCase,
	}

	err = handler.GetUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUseCase.AssertExpectations(t)

}
