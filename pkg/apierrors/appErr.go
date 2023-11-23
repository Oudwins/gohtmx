package apierrors

import (
	"errors"
	"net/http"
)

type AppErr struct {
	Status int
	err    error
}

type Unwrapper interface {
	Unwrap() error
}

var InternalServerError = New(errors.New("Something went wrong, please try again later"), http.StatusInternalServerError)

func New(err error, status int) error {
	newErr := &AppErr{
		Status: status,
		err:    err,
	}
	return newErr
}
func (e *AppErr) Error() string {
	return e.err.Error()
}

func (e *AppErr) Unwrap() error {

	var unwrapper Unwrapper
	if errors.As(e.err, &unwrapper) {
		return unwrapper.Unwrap()
	} else {
		return nil
	}

}
