package errs

import (
	"errors"
	"fmt"
)

var (
	ErrWrongInput          = New("wrong input")
	ErrNotFound            = New("entity not found")
	ErrSourceConnectionErr = New("connection error")
	ErrSourceInternal      = New("internal connection error")
)

type Error struct {
	err error
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v: %v", e.err.Error(), e.msg)
}

func (e *Error) Unwrap() error {
	return e.err
}

func Errf(err error, tmpl string, args ...interface{}) error {
	return &Error{err: err, msg: fmt.Sprintf(tmpl, args...)}
}

func New(text string) *Error {
	return &Error{err: errors.New(text)}
}