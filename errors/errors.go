package errors

import "errors"

// Custom error types
type Error struct {
	Code    int
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// Error constructors
func NewNotFoundError(message string) *Error {
	return &Error{
		Code:    404,
		Message: message,
	}
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Code:    400,
		Message: message,
	}
}

func NewInternalError(err error) *Error {
	return &Error{
		Code:    500,
		Message: "Internal server error",
		Err:     err,
	}
}

// Error checking functions to check error types
// Use it for special handling of not found errors
// Example:
//
//	if err := repository.FindUser(id); err != nil {
//	    if errors.IsNotFound(err) {
//	        // Special handling for not found case if needed
//	        return nil, fmt.Errorf("user not found: %w", err)
//	    }
//	    return nil, err
//	}
func IsNotFound(err error) bool {
	var e *Error
	return errors.As(err, &e) && e.Code == 404
}
