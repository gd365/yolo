package service

import "fmt"

type ErrorCode string

const (
	ErrCreateDir      ErrorCode = "E001"
	ErrCreatePkg      ErrorCode = "E002"
	ErrReadTemplate   ErrorCode = "E003"
	ErrParseTemplate  ErrorCode = "E004"
	ErrRenderTemplate ErrorCode = "E005"
	ErrWriteFile      ErrorCode = "E006"
	ErrGoMod          ErrorCode = "E007"
)

type YoloError struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func (e *YoloError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *YoloError) Unwrap() error {
	return e.Cause
}

func NewError(code ErrorCode, message string, cause error) *YoloError {
	return &YoloError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

var errorMessages = map[ErrorCode]string{
	ErrCreateDir:      "failed to create directory",
	ErrCreatePkg:      "failed to create package structure",
	ErrReadTemplate:   "failed to read template file",
	ErrParseTemplate:  "failed to parse template",
	ErrRenderTemplate: "failed to render template",
	ErrWriteFile:      "failed to write file",
	ErrGoMod:          "failed to initialize go module",
}

func ErrorWithCode(code ErrorCode, cause error) *YoloError {
	msg, ok := errorMessages[code]
	if !ok {
		msg = "unknown error"
	}
	return NewError(code, msg, cause)
}
