package models

import (
	"encoding/json"
	"fmt"
)

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type ErrorMessage struct {
	Msg string `json:"msg"`
}

type ErrorCause struct {
	ErrorType    string `json:"errorType"`
	ErrorMessage string `json:"errorMessage"`
}

func (e ErrorCode) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(`{"code": %v, "msg": "%s"}`, e.Code, e.Message)
	}
	return string(b)
}

func InvalidInput(error string) *ErrorCode {
	return &ErrorCode{
		Code:    1000,
		Message: fmt.Sprintf("invalid input. %v", error),
	}
}

func UnexpectedError(error string) *ErrorCode {
	return &ErrorCode{
		Code:    4000,
		Message: fmt.Sprintf("unexpected error. %v", error),
	}
}

func BadServerInitialization(error string) *ErrorCode {
	return &ErrorCode{
		Code:    4001,
		Message: fmt.Sprintf("bad server initialization. %v", error),
	}
}
