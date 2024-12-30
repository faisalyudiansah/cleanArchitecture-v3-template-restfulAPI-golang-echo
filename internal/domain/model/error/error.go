package domain_error

import "fmt"

type DomainError struct {
	HttpCode    int
	Code        string
	Message     string
	Err         error
	Description string `json:"description,omitempty"`
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
