package request

import (
	cErr "bookmark/internal/pkg/error"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Request interface {
	GetError(err error) *cErr.Error
}

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

var reg = regexp.MustCompile(`\[\d\]`)

func GetError(request interface{}, err error) *cErr.Error {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {

		_, isValidator := request.(Validator)

		for _, v := range err.(validator.ValidationErrors) {

			if isValidator {
				field := v.Field()
				field = reg.ReplaceAllString(field, ".*")

				if message, exist := request.(Validator).GetMessages()[field+"."+v.Tag()]; exist {
					return cErr.ValidateErr(message)
				}
			}

			return cErr.ValidateErr(v.Error())

		}

	}
	return cErr.ValidateErr("Parameter error")
}
