package utils

import (
	"net/http"

	httpError "dashboard-chatbot/bin/pkg/http-error"

	"github.com/labstack/echo/v4"
)

// Result common output
type Result struct {
	Data  interface{}
	Error interface{}
}

// BaseWrapperModel data structure
type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}

// Response function
func Response(data interface{}, message string, code int, c echo.Context) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return c.JSON(code, result)
}

// ErrorStatusCode function
func getErrorStatusCode(err interface{}) httpError.CommonError {
	errData := httpError.CommonError{}

	switch obj := err.(type) {
	case httpError.BadRequest:
		errData.ResponseCode = http.StatusBadRequest
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Unauthorized:
		errData.ResponseCode = http.StatusUnauthorized
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Conflict:
		errData.ResponseCode = http.StatusConflict
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.NotFound:
		errData.ResponseCode = http.StatusNotFound
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.InternalServerError:
		errData.ResponseCode = http.StatusInternalServerError
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	default:
		errData.Code = http.StatusConflict
		return errData
	}
}

// ResponseError function
func ResponseError(err interface{}, c echo.Context) error {
	errObj := getErrorStatusCode(err)
	result := BaseWrapperModel{
		Success: false,
		Data:    errObj.Data,
		Message: errObj.Message,
		Code:    errObj.Code,
	}

	return c.JSON(errObj.ResponseCode, result)
}
