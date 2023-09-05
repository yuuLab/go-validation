package validation

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/yuuLab/go-validation/presentation/response"
)

// ozzoValidator implements the binding.StructValidator.
type ozzoValidator struct {
	validator *Validator
}

func NewOzzoValidator() binding.StructValidator {
	return &ozzoValidator{validator: &Validator{}}
}

func (v *ozzoValidator) ValidateStruct(obj any) error {
	return v.validator.Validate(obj)
}

func (v *ozzoValidator) Engine() any {
	return v.validator
}

// Validator performs validation on parameters.
type Validator struct{}

// Validate performs validation on any given interface.
// If the provided interface is valid or doesn't implement the validation.Validatable interface, it returns nil.
func (v Validator) Validate(obj any) error {
	if obj == nil {
		return nil
	}

	val, ok := obj.(validation.Validatable)
	if !ok {
		return nil
	}
	if err := val.Validate(); err != nil {
		if verr, ok := err.(validation.Errors); ok {
			var params []response.InvalidParams
			for key, val := range verr {
				params = append(params, response.InvalidParams{Name: key, Reason: val.Error()})
			}
			return v.newValidationError(params, err)
		}
		return v.newServerError(err)
	}
	return nil
}

// newValidationError creates an error that wraps a Internal Error.
func (v Validator) newServerError(err error) error {
	return ValidationError{
		response: response.ValidationError{Type: "SERVER_ERROR", Title: "unexpected errors"},
		status:   http.StatusInternalServerError,
		err:      err,
	}
}

// newValidationError creates an error that wraps a Validation Error.
func (v Validator) newValidationError(params []response.InvalidParams, err error) error {
	return ValidationError{
		response: response.ValidationError{Type: "VALIDATION_ERROR", Title: "Your request parameters didn't validate.", Pramas: params},
		status:   http.StatusBadRequest,
		err:      err,
	}
}
