package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"timesheet-app/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
)

func GetValidationError(err error) []response.ValidationField {

	var validationField []response.ValidationField
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, validationError := range ve {
			log.Debug().Msg(fmt.Sprintf("validationError: %v", validationError))
			myField := convertFieldRequired(validationError.Namespace())
			validationField = append(validationField, response.ValidationField{
				FieldName: myField,
				Message:   formatMessage(validationError),
			})
		}
	}
	return validationField
}

func convertFieldRequired(myValue string) string {

	log.Debug().Msg("convertFieldRequired: " + myValue)
	fieldSegment := strings.Split(myValue, ".")

	myField := ""
	length := len(fieldSegment)
	i := 1
	for _, val := range fieldSegment {
		if i == 1 {
			i++
			continue
		}

		if i == length {
			myField += strcase.SnakeCase(val)
			break
		}

		myField += strcase.LowerCamelCase(val) + `/`
		i++
	}
	return myField
}

func formatMessage(err validator.FieldError) string {

	var message string

	switch err.Tag() {
	case "required":
		message = "required"
	case "number":
		message = "must be number"
	case "email":
		message = "invalid format email"
	case "DateOnly":
		message = "invalid format date"
	case "min":
		message = "minimum value is not exceed"
	case "max":
		message = "maximum value is exceed"
	case "password":
		message = "password must contain at least one uppercase letter, one lowercase letter, one number and one special character"
	case "nomorHp":
		message = "invalid number phone format"
	}
	return message
}

func ValidatePassword(fl validator.FieldLevel) bool {

	password := fl.Field().String()

	uppercase := false
	lowercase := false
	number := false
	specialChar := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			uppercase = true
		case 'a' <= char && char <= 'z':
			lowercase = true
		case '0' <= char && char <= '9':
			number = true
		case strings.ContainsRune("!@#$%^&*()-_=+{}[]|;:,.<>?/~", char):
			specialChar = true
		}
	}
	return uppercase && lowercase && number && specialChar
}

func ValidateNoHp(fl validator.FieldLevel) bool {
	noHp := fl.Field().String()
	return regexp.MustCompile(`^(08|\+62)\d{8,20}$`).MatchString(noHp)
}
