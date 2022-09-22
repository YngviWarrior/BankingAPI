package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (v Validator) InputValidator(obj any) []string {
	valid := validator.New()

	err := valid.Struct(obj)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var tag string
			switch e.Tag() {
			case "min", "max":
				tag = fmt.Sprintf(": check %v length ", e.StructField())
			}

			s := fmt.Sprintf("Input field %s is invalid%v", e.StructField(), tag)
			v.Error = append(v.Error, s)
		}

		return v.Error
	}

	return v.Error
}
