package utils

import "net/http"

type AppErrorInterface interface {
	Error() string
}

type AppError struct {
	Message    string
	StatusCode int
	Success    bool
}

func (appError *AppError) Error() string {
	return appError.Message
}

func NotFoundError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Success:    false,
	}
}

func UnauthorizedError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Success:    false,
	}
}

func InternalServerError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Success:    false,
	}
}

func BadRequestError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}
}

func ForbiddenError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Success:    false,
	}
}
