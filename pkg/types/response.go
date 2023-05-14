package types

import "github.com/pangud/pangud/pkg/errors"

// Response is the response of the API
type ResponseEntity[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// NewResponseEntity create a new response entity
func NewResponseEntity[T any](err errors.Error, data T) *ResponseEntity[T] {
	return &ResponseEntity[T]{Code: int(err.Code), Msg: err.Message, Data: data}
}
