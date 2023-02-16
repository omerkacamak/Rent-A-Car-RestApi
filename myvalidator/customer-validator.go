package myvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/omerkacamak/rentacar-golang/entity"
)

func CustomerValidator(field validator.FieldLevel) bool {
	result := LengthValidator(field)

	return result
}

func LengthValidator(field validator.FieldLevel) bool {
	if field.Field().Len() < 5 {
		// err := errors.New("Karakter uzunluğu 5ten küçük olamaz")
		return false
	}
	return true
}

func CustomerValidator2(validate *validator.Validate, customer *entity.Customer) error {
	validate = validator.New()
	validate.RegisterValidation("lent", CustomerValidator)
	validate.RegisterValidation("deneme", CustomerValidator)
	err := validate.Struct(&customer)
	if err != nil {
		println("asdasd::::  " + err.Error())
		return err
	}
	return nil

}
