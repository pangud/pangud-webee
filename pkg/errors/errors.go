package errors

import (
	"errors"
	"fmt"
)

type ErrorCode uint32

const (
	ErrCodeOK           ErrorCode = 0
	ErrCodeUnknownError ErrorCode = 99999

	// 以1开头为系统端错误 1  00 标识 通用错误
	ErrCodeSystemError ErrorCode = 100000
	// 以1开头为系统端错误 1  00 标识 通用错误 数据库错误
	ErrCodeDBError           ErrorCode = 100010
	ErrCodeDBConnectionError ErrorCode = 100011

	ErrCodeUserError    ErrorCode = 200000
	ErrCodeInvalidParam ErrorCode = 200001 //参数错误
	// 以3开头为外部系统错误 3  00 标识通用错误
	ErrCodeExternalError ErrorCode = 300000
)

var (
	ok                = New(ErrCodeOK, "success")
	SystemError       = New(ErrCodeSystemError, "system internal server error")
	DBError           = New(ErrCodeDBError, "database error")
	DBConnectionError = New(ErrCodeDBError, "database connection error")
	UserError         = New(ErrCodeUserError, "user error")
	ExternalError     = New(ErrCodeUserError, "external system error")
)

func OK() Error {
	return *ok
}

type Error struct {
	Code    ErrorCode `json:"code" example:"0"`
	Message string    `json:"msg" example:"success"`
	cause   error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d message = %s cause = %v", e.Code, e.Message, e.cause)
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Unwrap() error { return e.cause }

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Code == e.Code
	}
	return false
}

// WithCause with the underlying cause of the error.
func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

// Clone deep clone error to a new error.
func Clone(err *Error) *Error {
	if err == nil {
		return nil
	}

	return &Error{
		cause:   err.cause,
		Code:    err.Code,
		Message: err.Message,
	}
}

func New(code ErrorCode, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}

//func OK() *Error {
//	return New(ErrCodeOK, "success")
//}

func FromError(err error) (s *Error, ok bool) {
	if err == nil {
		return nil, true
	}
	if se, ok := err.(*Error); ok {
		return se, true
	}
	return New(ErrCodeUnknownError, err.Error()), false
}
