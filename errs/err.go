package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (app AppError) Error() string {
	return app.Message
}

func ErrOperatorNotFound() error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: "Operator Not Found",
	}
}
