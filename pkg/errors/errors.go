package errors

const (
	CodeOk                = 0    // OK
	CodeErr               = -1   // general error
	CodeAuthErr           = -2   // unauthenticated request
	CodeTokenErr          = -3   // unauthenticated request
	CodeBadRequestErr     = 400 // bad request
	CodeForbiddenErr      = 403 // forbidden request
	CodeNotFoundErr       = 404 // not found
	CodeInternalErrorErr  = 500 // bad request
	CodeNotImplementedErr = 501 // bad request

)

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New(CodeTokenErr, "Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New(CodeAuthErr, "Unauthorized")

	// ErrBadRequest is returned when a resource is bad request.
	ErrBadRequest = New(CodeBadRequestErr, "Bad Request")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New(CodeForbiddenErr, "Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New(CodeNotFoundErr, "Not Found")

	// ErrInternalError is returned when an endpoint is internal error.
	ErrInternalError = New(CodeInternalErrorErr, "Internal Error")

	// ErrNotImplemented is returned when an endpoint is not implemented.
	ErrNotImplemented = New(CodeNotImplementedErr, "Not Implemented")
)

// Error represents a json-encoded API error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(code int, text string) error {
	return &Error{
		Code:    code,
		Message: text,
	}
}
