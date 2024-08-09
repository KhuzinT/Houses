package utils

import "errors"

// repo errors

var ErrUserNotFound = errors.New("user not found")
var ErrFlatNotFound = errors.New("flat not found")
var ErrHouseNotFound = errors.New("house not found")

// login errors

var ErrPasswordNotMatch = errors.New("password not match")
var ErrPasswordTooShort = errors.New("password too short")
var ErrEmptyEmail = errors.New("empty email")

// auth errors

var ErrMissingAuthHeader = errors.New("missing auth header")

// handler errors

var ErrUnknownFlatStatus = errors.New("unknown flat status")
