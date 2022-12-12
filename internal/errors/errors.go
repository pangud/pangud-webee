package errors

import (
	"errors"
	"fmt"
)

type ErrorCode uint32

const (
	ErrCodeOK ErrorCode = iota

	/*	 以1开头为用户错误
	11000-11099 用户通用错误*/

	ErrCodeUserError    ErrorCode = 11000
	ErrCodeInvalidParam ErrorCode = 11001
	ErrCodeNotFound     ErrorCode = 11002
	ErrCodeUnknownError ErrorCode = 11099

	// 以2开头为系统端错误
	//21000-21099 系统通用错误

	ErrCodeSystemError ErrorCode = 20000
	ErrcodeDBError     ErrorCode = 20001
)

var (
	InternalServer = New(ErrCodeSystemError, "internal server error")
	BadRequest     = New(ErrCodeInvalidParam, "invalid param")
	NotFound       = New(ErrCodeNotFound, "not found")
	OK             = New(ErrCodeOK, "success")
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"msg"`
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
