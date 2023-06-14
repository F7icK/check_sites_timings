package customerr

import "net/http"

type CustomError struct {
	msg  string
	Code int
}

func NewError(msg string, code int) *CustomError {
	return &CustomError{
		msg:  msg,
		Code: code,
	}
}

func (c *CustomError) Error() string {
	return c.msg
}

var (
	ErrorInternalServerError = NewError(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	ErrorBadRequest          = NewError(http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
)
