package validation

import "github.com/yuuLab/go-validation/presentation/response"

// ValidationError represents validation error.
type ValidationError struct {
	response response.ValidationError
	status   int
	err      error
}

// Error returns the error string of Errors.
func (v ValidationError) Error() string {
	return v.err.Error()
}

// Response returns the Response.
func (v ValidationError) Response() response.ValidationError {
	return v.response
}

// Status returns the Status Code.
func (v ValidationError) Status() int {
	return v.status
}
