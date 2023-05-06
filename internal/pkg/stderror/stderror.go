package stderror

import (
	"errors"
	"fmt"
)

type StdError struct {
	HttpCode     int    `json:"-"`
	InternalCode string `json:"internalCode"`
	Error        error  `json:"-"`
	Message      string `json:"message,omitempty"`
}

func New(code int, internalCode string, errMsg ...any) StdError {
	var format string
	for i, _ := range errMsg {
		if i == 0 {
			format = format + "%s"
		} else {
			format = format + " %s"
		}
	}

	return StdError{
		HttpCode:     code,
		InternalCode: internalCode,
		Error:        fmt.Errorf(format, errMsg...),
	}
}

func (err *StdError) ErrorMsg(message error) *StdError {
	err.Error = errors.New(message.Error())
	return err
}
