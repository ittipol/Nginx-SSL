package apps

import (
	"go-nginx-ssl/errs"

	v "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidatePayload(req interface{}) error
}

type validator struct {
	validate *v.Validate
}

func Newvalidator(validate *v.Validate) Validator {
	return &validator{validate}
}

func (obj validator) ValidatePayload(req interface{}) error {
	if err := obj.validate.Struct(req); err != nil {
		return errs.NewBadRequestError()
	}
	return nil
}
