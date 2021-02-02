package errorhandler

import (
	"net/http"
)

type ErrorType string

func (ec ErrorType) String() string {
	return string(ec)
}

const (
	ErrBadParams       ErrorType = "bad_params"
	ErrEmptyBody       ErrorType = "empty_body"
	ErrInvalidRequest  ErrorType = "invalid_request"
	ErrUnauthorized    ErrorType = "unauthorized"
	ErrNotFound        ErrorType = "not_found"
	ErrUniqueViolation ErrorType = "unique_violation"
	ErrDatabase        ErrorType = "database_error"
	ErrInternal        ErrorType = "internal_error"
	ErrExternalAPI     ErrorType = "external_api_error"
	ErrForbidden       ErrorType = "forbidden"
	ErrUnknown         ErrorType = "unknown_error"
)

var pgErrorTypeMap = map[ErrorType]string{
	ErrUniqueViolation: "23505",
}

var codeStatusMap = map[ErrorType]int{
	ErrBadParams:       http.StatusBadRequest,
	ErrEmptyBody:       http.StatusBadRequest,
	ErrInvalidRequest:  http.StatusBadRequest,
	ErrUnauthorized:    http.StatusUnauthorized,
	ErrNotFound:        http.StatusNotFound,
	ErrUniqueViolation: http.StatusConflict,
	ErrDatabase:        http.StatusInternalServerError,
	ErrInternal:        http.StatusInternalServerError,
	ErrExternalAPI:     http.StatusInternalServerError,
	ErrForbidden:       http.StatusForbidden,
	ErrUnknown:         http.StatusInternalServerError,
}

func GetHTTPStatus(code ErrorType) int {
	return codeStatusMap[code]
}

func GetPgErrorType(code ErrorType) string {
	return pgErrorTypeMap[code]
}
