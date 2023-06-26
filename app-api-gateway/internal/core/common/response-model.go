package common

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

var MissingData = errors.New("missing data")
var InvalidData = errors.New("invalid data")

type HttpReponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func ResponseFromError(err error) (int, HttpReponse) {
	var code int = -1
	var message string = "Bad request"

	var errorCode = status.Code(err)

	switch errorCode {
	case codes.NotFound:
		code = http.StatusNotFound
		message = "We looked everywhere but couldn't find what you were looking for"
		break
	case codes.InvalidArgument:
		code = http.StatusBadRequest
		message = "Your message is probably badly formatted"
		break
	case codes.AlreadyExists:
		code = http.StatusConflict
		message = "The resource you are trying to create already exists"
		break
	case codes.PermissionDenied:
		code = http.StatusForbidden
		message = "You don't have permission to do that"
		break
	default:
		log.Printf("Error: %v", err)
		log.Printf("Error code: %v", errorCode)
		code = 500
		message = "We can't figure out what went wrong"
		break
	}

	return code, HttpReponse{
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  []string{status.Convert(err).Message()},
	}
}

func ResponseFromErrorWithDetails(code codes.Code, message string) (int, HttpReponse) {
	return ResponseFromError(status.Error(code, message))
}
