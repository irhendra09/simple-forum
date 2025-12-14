package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrConflict     = errors.New("conflict")
	ErrBadRequest   = errors.New("bad request")
)

// ToHTTPStatus maps application errors to HTTP status codes.
func ToHTTPStatus(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrConflict:
		return http.StatusConflict
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
