package error

import "net/http"

const (
	DEFAULT_ERROR     = 400
	NOT_FOUND         = 404
	TOKEN_ERROR       = 401
	FORBIDDEN         = 403
	VALIDATE_ERROR    = 422
	TOO_MANY_REQUESTS = 429
	SERVER_ERROR      = 500
)

type Error struct {
	httpCode  int
	errorCode int
	errorMsg  string
}

func New(httpCode int, errorCode int, errorMsg string) *Error {
	return &Error{
		httpCode:  httpCode,
		errorCode: errorCode,
		errorMsg:  errorMsg,
	}
}

func BadRequest(errorMsg string, errorCode ...int) *Error {
	errCode := DEFAULT_ERROR
	if len(errorCode) > 0 {
		errCode = errorCode[0]
	}
	return New(http.StatusBadRequest, errCode, errorMsg)
}

func Unauthorized(errorMsg string, errorCode ...int) *Error {
	errCode := TOKEN_ERROR
	if len(errorCode) > 0 {
		errCode = errorCode[0]
	}
	return New(http.StatusUnauthorized, errCode, errorMsg)
}

func Forbidden(errorMsg string, errorCode ...int) *Error {
	errCode := FORBIDDEN
	if len(errorCode) > 0 {
		errCode = errorCode[0]
	}
	return New(http.StatusForbidden, errCode, errorMsg)
}

func NotFound(errorMsg string, errorCode ...int) *Error {
	errCode := NOT_FOUND
	if len(errorCode) > 0 {
		errCode = errorCode[0]
	}
	return New(http.StatusNotFound, errCode, errorMsg)
}

func ValidateErr(errorMsg string) *Error {
	return New(http.StatusUnprocessableEntity, VALIDATE_ERROR, errorMsg)
}

func TooManyRequestsErr(errorMsg string) *Error {
	return New(http.StatusTooManyRequests, TOO_MANY_REQUESTS, errorMsg)
}

func InternalServerError(errorMsg string) *Error {
	return New(http.StatusInternalServerError, SERVER_ERROR, errorMsg)
}

func (e *Error) HttpCode() int {
	return e.httpCode
}

func (e *Error) ErrorCode() int {
	return e.errorCode
}

func (e *Error) Error() string {
	return e.errorMsg
}
