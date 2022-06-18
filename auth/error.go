package auth

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserEmailAlreadyExists = errors.New("user email already exists")
	ErrUserPhoneAlreadyExists = errors.New("user phone already exists")
	ErrUserEmailNotFound = errors.New("user email not found")
	ErrUserPhoneNotFound = errors.New("user phone not found")
	ErrUserEmailNotConfirmed = errors.New("user email not confirmed")
	ErrUserPhoneNotConfirmed = errors.New("user phone not confirmed")
	ErrUserPasswordNotMatch = errors.New("user password not match")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrCreatingUser = errors.New("error creating user")
)