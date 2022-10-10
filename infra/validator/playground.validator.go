package validator

import (
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator"
)

type Date struct {
	DateField string `validate:"date"`
}

func (v *Validator) InputValidator(obj any) []string {
	valid := validator.New()

	valid.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		date := fl.Field()
		_, err := time.Parse("2006-01-02", date.String())
		if err != nil {
			log.Println(err)
			return false
		}

		return true
	}, false)

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
