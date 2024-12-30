package domain_error

import "net/http"

type MandatoryFieldError struct {
	DomainError
	Description string `json:"description,omitempty"`
}

func NewMandatoryFieldError(description string) error {
	e := &MandatoryFieldError{}
	e.HttpCode = http.StatusBadRequest
	e.Code = "402004"
	e.Message = "mandatory field is empty"
	e.Description = description
	return e
}
