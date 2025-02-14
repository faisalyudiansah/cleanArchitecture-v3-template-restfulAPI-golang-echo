package httpresponse

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoResponse struct{}

type ResponseData struct {
	Data string `json:"data"`
}

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Response
	Error interface{} `json:"error,omitempty"`
}

type DataResponse[T any] struct {
	Response
	Data   T            `json:"data,omitempty"`
	Errors []FieldError `json:"errors,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type PaginatedReponse[T any] struct {
	Data        []T   `json:"data"`
	Total       int64 `json:"total"`
	CurrentPage int   `json:"current_page,omitempty"`
	PerPage     int   `json:"per_page,omitempty"`
	Page        int   `json:"page,omitempty"`
}

func (r *PaginatedReponse[T]) EchoJsonResponse(c echo.Context) error {
	return c.JSON(http.StatusOK, r)
}

func (r *DataResponse[T]) EchoJsonResponse(c echo.Context) error {
	return c.JSON(http.StatusOK, r)
}

func (r *DataResponse[T]) EchoJsonResponseWithCodeMessage(c echo.Context) error {
	return c.JSON(r.Code, r)
}

func (r *ErrorResponse) EchoJsonResponse(c echo.Context) error {
	return c.JSON(r.Code, r)
}

func (r *Response) EchoJsonResponse(c echo.Context) error {
	return c.JSON(r.Code, r)
}
