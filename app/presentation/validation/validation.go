package validation

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/yuuLab/go-validation/presentation/response"
)

// ValidationError represents validation error.
type ValidationError struct {
	response response.ValidationError
	status   int
	err      error
}

// Validator performs validation on parameters.
type Validator struct{}

// NewValidator creates a new validator.
func NewValidator() *Validator {
	return &Validator{}
}

// Validate performs validation on any given interface.
// If the provided interface is valid or doesn't implement the validation.Validatable interface, it returns nil.
func (v Validator) Validate(i interface{}) error {
	val, ok := i.(validation.Validatable)
	if !ok {
		return nil
	}
	if err := val.Validate(); err != nil {
		if verr, ok := err.(validation.Errors); ok {
			var params []response.InvalidParams
			for key, val := range verr {
				params = append(params, response.InvalidParams{Name: key, Reason: val.Error()})
			}
			return newValidationError(params, err)
		}
		return newServerError(err)
	}
	return nil
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

// newValidationError creates an error that wraps a Internal Error.
func newServerError(err error) error {
	return ValidationError{
		response: response.ValidationError{Type: "SERVER_ERROR", Title: "unexpected errors"},
		status:   http.StatusInternalServerError,
		err:      err,
	}
}

// newValidationError creates an error that wraps a Validation Error.
func newValidationError(params []response.InvalidParams, err error) error {
	return ValidationError{
		response: response.ValidationError{Type: "VALIDATION_ERROR", Title: "validation errors", Pramas: params},
		status:   http.StatusBadRequest,
		err:      err,
	}
}
