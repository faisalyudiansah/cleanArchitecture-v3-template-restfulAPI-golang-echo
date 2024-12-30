package httpmiddleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/internal/adapter/validator/tools"
	httpresponse "server/pkg/common/http/response"
	"server/pkg/constant"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	var code int
	var response httpresponse.DataResponse[any]

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		handleValidationError(c, validationErrors)
		return
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		handleUnmarshalTypeError(c, unmarshalTypeError)
		return
	}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if msg, ok := he.Message.(string); ok {
			response = httpresponse.DataResponse[any]{
				Response: httpresponse.Response{
					Code:    code,
					Message: msg,
				},
			}
		} else if msg, ok := he.Message.(httpresponse.DataResponse[any]); ok {
			response = msg
		} else {
			response = httpresponse.DataResponse[any]{
				Response: httpresponse.Response{
					Code:    code,
					Message: constant.CurrentlyUnexpectedError,
				},
			}
		}
		c.JSON(code, response)
		return
	}

	response = httpresponse.DataResponse[any]{
		Response: httpresponse.Response{
			Code:    http.StatusInternalServerError,
			Message: constant.UnexpectedErrorOccurred,
		},
	}
	c.JSON(http.StatusInternalServerError, response)
}

func handleValidationError(c echo.Context, err validator.ValidationErrors) error {
	var errors []httpresponse.FieldError
	for _, vErr := range err {
		field := vErr.Field()
		// tag := vErr.Tag()
		errors = append(errors, httpresponse.FieldError{
			Field:   field,
			Message: tools.TagToMsg(vErr),
		})
	}
	return c.JSON(http.StatusBadRequest, httpresponse.DataResponse[any]{
		Response: httpresponse.Response{
			Code: http.StatusUnprocessableEntity,
		},
		Errors: errors,
	})
}

func handleUnmarshalTypeError(c echo.Context, err *json.UnmarshalTypeError) error {
	response := httpresponse.DataResponse[any]{
		Response: httpresponse.Response{
			Code: http.StatusUnprocessableEntity,
		},
		Errors: []httpresponse.FieldError{
			{
				Field:   err.Field,
				Message: constant.MismatchInput,
			},
		},
	}
	return c.JSON(http.StatusUnprocessableEntity, response)
}
