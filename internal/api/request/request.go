package request

import (
	cErr "bookmark/internal/pkg/error"
	"regexp"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
					return cErr.New(http.StatusOK, 422, message)
				}
			}

			return cErr.New(http.StatusOK, 422, v.Error())

		}

	}
	return cErr.New(http.StatusOK, 422, "Parameter error")
}

func ShouldBindWith(c *gin.Context, obj any, b binding.Binding) error {
	err := c.ShouldBindWith(obj, b)
	if err != nil {
		return GetError(obj, err)
	}
	return nil
}

func ShouldBindWithQuery(c *gin.Context, obj any) error {
	return ShouldBindWith(c, obj, binding.Query)
}

func ShouldBindWithJSON(c *gin.Context, obj any) error {
	return ShouldBindWith(c, obj, binding.JSON)
}
