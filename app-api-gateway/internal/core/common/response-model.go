package common

import "errors"

var NotFoundErr = errors.New("no rows in result set")
var MissingData = errors.New("missing data")
var InvalidData = errors.New("invalid data")

type HttpReponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func ResponseFromError(err error) (int, HttpReponse) {
	var code int = 400
	var message string = "Bad request"

	switch {
	case errors.Is(err, NotFoundErr):
		code = 404
		message = "Not found"
		break
	case errors.Is(err, MissingData):
		code = 400
		message = "Invalid request"
		break
	case errors.Is(err, InvalidData):
		code = 400
		message = "Invalid request"
		break
	default:
		code = 500
		message = "Internal server error"
		break
	}

	return code, HttpReponse{
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  []string{err.Error()},
	}
}
