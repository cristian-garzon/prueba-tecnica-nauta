package model

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorCode string

const (
	ErrNoBookingsFound   ErrorCode = "NOT_BOOKINGS_FOUND"
	ErrNoBookingFound    ErrorCode = "NOT_BOOKING_FOUND"
	ErrNotClientFound    ErrorCode = "NOT_CLIENT_FOUND"
	ErrNoOrdersFound     ErrorCode = "NOT_ORDERS_FOUND"
	ErrNoContainersFound ErrorCode = "NOT_CONTAINERS_FOUND"

	ErrInvalidClientId    ErrorCode = "INVALID_CLIENT_ID"
	ErrInvalidBookingId   ErrorCode = "INVALID_BOOKING_ID"
	ErrInvalidOrderId     ErrorCode = "INVALID_ORDER_ID"
	ErrInvalidContainerId ErrorCode = "INVALID_CONTAINER_ID"
	ErrInvalidEmail       ErrorCode = "INVALID_EMAIL"

	InvalidBodyError ErrorCode = "INVALID_BODY"

	ErrQueryError ErrorCode = "QUERY_ERROR"
	UnknownError  ErrorCode = "UNKNOWN_ERROR"
)

func mapErrorCode(code ErrorCode) int {
	switch code {
	case ErrNoBookingFound, ErrNotClientFound, ErrNoContainersFound, ErrNoOrdersFound:
		return http.StatusNotFound
	case ErrInvalidClientId, ErrInvalidBookingId, ErrInvalidOrderId, ErrInvalidContainerId, ErrInvalidEmail, InvalidBodyError:
		return http.StatusBadRequest
	case ErrQueryError, UnknownError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var statusCode int
	var response AppError

	if appErr, ok := err.(*AppError); ok {
		statusCode = appErr.Code
		response = AppError{
			Message:   appErr.Message,
			Code:      appErr.Code,
			ErrorCode: appErr.ErrorCode,
		}
	} else {
		statusCode = http.StatusInternalServerError
		response = AppError{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: UnknownError,
		}
	}

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding error response: %v", err)
	}
}

type AppError struct {
	Message   string
	Code      int
	ErrorCode ErrorCode
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code ErrorCode) *AppError {
	return &AppError{Message: message, Code: mapErrorCode(code), ErrorCode: code}
}
