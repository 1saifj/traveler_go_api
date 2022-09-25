package utils

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	code int
	err  error
	msg  string
}

func NewError(code int, err error, message string) *Error {
	return &Error{
		code: code,
		err:  err,
		msg:  message,
	}
}

func NewErrorf(code int, err error, format string, args ...interface{}) *Error {
	return &Error{
		code: code,
		err:  errors.New(fmt.Sprintf(format, args...)),
	}
}

func NewBadRequestError(err error) *Error {
	return NewError(fiber.StatusBadRequest, err, "bad request")
}

func NewUnauthorizedError(err error) *Error {
	return NewError(fiber.StatusUnauthorized, err, "unauthorized")
}

func NewForbiddenError(err error) *Error {
	return NewError(fiber.StatusForbidden, err, "forbidden")
}

func NewNotFoundError(err error) *Error {
	return NewError(fiber.StatusNotFound, err, "not found")
}

func (e Error) Code() int {
	if e.code == 0 {
		return fiber.StatusInternalServerError
	}

	return e.code
}

func (e Error) Error() string {
	return e.err.Error()
}
