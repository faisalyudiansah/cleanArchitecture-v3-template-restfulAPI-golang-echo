package constant

const (
	EntityNotFoundErrorMessage        = "%s not found"
	ErrValidateRequest                = "failed to validate request"
	ErrDecodeFilterRequest            = "failed to decode filter request"
	RequestTimeoutErrorMessage        = "failed to process request in time, please try again"
	UnauthorizedErrorMessage          = "unauthorized"
	EOFErrorMessage                   = "missing body request"
	StrConvSyntaxErrorMessage         = "invalid syntax for %s"
	ForbiddenAccessErrorMessage       = "you do not have permission to access this resource"
	TooManyRequestsErrorMessage       = "the server is experiencing high load, please try again later"
	InternalServerErrorMessage        = "internal server error"
	ResetPasswordErrorMessage         = "please try again later"
	ValidationErrorMessage            = "input validation error"
	InvalidJsonUnmarshallErrorMessage = "invalid JSON format"
	JsonSyntaxErrorMessage            = "invalid JSON syntax"
	InvalidJsonValueTypeErrorMessage  = "invalid value for %s"
	InvalidIDErrorMessage             = "expected a numeric value"
	InvalidRequestBody                = "invalid request body"
	UnexpectedErrorOccurred           = "unexpected error occurred"
	MismatchInput                     = "type mismatch, check your input format"
	CurrentlyUnexpectedError          = "currently our server is facing unexpected error, please try again later"
)
