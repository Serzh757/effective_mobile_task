package errors

// Error response json object.
type Error struct {
	Message string `json:"message"` // Description of the error.
}

// newErrorByWrappingError wraps another error message.
func newErrorByWrappingError(err error) interface{} {
	return &Error{
		Message: err.Error(),
	}
}
