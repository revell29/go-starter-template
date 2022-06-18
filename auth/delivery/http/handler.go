package http

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/questizen/core-system/domain"
	"github.com/questizen/core-system/helpers"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

// AuthHandler represent the httphandler for auth
type Handler struct {
	AUsecase domain.AuthUseCase
}

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type ResponseGetUser struct {
	Data []domain.User `json:"data"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewAuthHandler will initialize the auth/ resources endpoint
func NewAuthHandler(e *echo.Echo, usecase domain.AuthUseCase) {
	handler := &Handler{
		AUsecase: usecase,
	}

	// r.HandleFunc("/signup", handler.CreateUser).Methods("POST")
	e.POST("/signup", handler.CreateUser)
	e.GET("/users", handler.GetUser)
}

// isRequestValid will validate the request body
func isRequestValid(m *domain.AuthUser) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Create user handler
func (h *Handler) CreateUser(c echo.Context) (err error) {
	var input domain.AuthUser
	err = c.Bind(&input)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&input); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(password)

	data, err := h.AUsecase.CreateUser(ctx, &input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) GetUser(c echo.Context) (err error) {
	// numS := c.QueryParam("num")
	// num, _ := strconv.Atoi(numS)
	// cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listAr, err := h.AUsecase.GetUser(ctx)
	if err != nil {
		return c.JSON(helpers.GetStausCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseGetUser{Data: listAr})
}
