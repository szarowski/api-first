package errors

import (
	"errors"
	"net/http"
	"strict-server/internal/api/gen"
)

// Internal error codes
const (
	ErrorInternalServerError      = "INTERNAL_SERVER_ERROR"
	ErrorInvalidUserRequestParams = "INVALID_USER_REQUEST_PARAMS"
	ErrorInvalidUserRequestBody   = "INVALID_USER_REQUEST_BODY"
	ErrorStatusConflictUser       = "CONFLICT_USER"
	ErrorStatusNotFoundUser       = "NOT_FOUND_USER"
	ErrorStatusPageNotFound       = "PAGE_NOT_FOUND"
	ErrorStatusMethodNotAllowed   = "METHOD_NOT_ALLOWED"
)

// Internal errors
var (
	ErrInternalServerError    = errors.New(ErrorInternalServerError)
	ErrInvalidUserInputData   = errors.New(ErrorInvalidUserRequestParams)
	ErrStatusConflictUser     = errors.New(ErrorStatusConflictUser)
	ErrStatusNotFoundUser     = errors.New(ErrorStatusNotFoundUser)
	ErrStatusPageNotFound     = errors.New(ErrorStatusPageNotFound)
	ErrStatusMethodNotAllowed = errors.New(ErrorStatusMethodNotAllowed)
)

// Errors - A map of ErrorResult representations of the errors
var Errors = map[error]*gen.ErrorResult{
	ErrInternalServerError: {
		Code:       ErrorInternalServerError,
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	},
	ErrInvalidUserInputData: {
		Code:       ErrorInvalidUserRequestParams,
		Message:    "Invalid user input data",
		StatusCode: http.StatusBadRequest,
	},
	ErrStatusConflictUser: {
		Code:       ErrorStatusConflictUser,
		Message:    "User id in conflict",
		StatusCode: http.StatusConflict,
	},
	ErrStatusNotFoundUser: {
		Code:       ErrorStatusNotFoundUser,
		Message:    "User not found",
		StatusCode: http.StatusNotFound,
	},
	ErrStatusPageNotFound: {
		Code:       ErrorStatusPageNotFound,
		Message:    "Page not found",
		StatusCode: http.StatusNotFound,
	},
	ErrStatusMethodNotAllowed: {
		Code:       ErrorStatusMethodNotAllowed,
		Message:    "Method not allowed",
		StatusCode: http.StatusMethodNotAllowed,
	},
}
