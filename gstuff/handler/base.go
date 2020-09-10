package handler

import (
	"net/http"

	"log"

	"github.com/kenkinsai/ntc-common/config"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

var cfg = config.GetConfig()

// NewValidator ..
func NewValidator() *MyValidator {
	return &MyValidator{validator: validator.New()}
}

// MyValidator ..
type MyValidator struct {
	validator *validator.Validate
}

// Validate ..
func (cv *MyValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// GetRequestID ..
func GetRequestID(c echo.Context) (requestID string) {
	if c.Get("reqID") != nil {
		requestID = c.Get("reqID").(string)
	}
	return
}

// Health func
func Health(c echo.Context) (err error) {
	return c.JSON(Success(nil))
}

// ResponseContent struct
type ResponseContent struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error func
func Error(err error, c echo.Context) {
	code := http.StatusBadRequest
	msg := http.StatusText(code)

	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
		msg = httpError.Message.(string)
	}

	if cfg.Debug {
		msg = err.Error()
	}

	log.Println(err)

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			c.NoContent(code)
		} else {
			c.JSON(code, &ResponseContent{Code: code, Message: msg})
		}
	}
}

// Success func
func Success(data interface{}) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

// NotFound func
func NotFound(message string) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
