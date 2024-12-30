package tools

import (
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

var (
	uppercaseRegexp   = regexp.MustCompile(`[A-Z]`)
	numberRegexp      = regexp.MustCompile(`[0-9]`)
	specialCharRegexp = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':",\.<>\/\\\|` + "`" + `~]`)
	phoneRegexp       = regexp.MustCompile(`^(\+62|62|08)[0-9]{8,12}$`)
	dayOfWeeks        = map[string]bool{
		"Sunday":    true,
		"Monday":    true,
		"Tuesday":   true,
		"Wednesday": true,
		"Thursday":  true,
		"Friday":    true,
		"Saturday":  true,
	}
)

func ValidateExcluded(fl validator.FieldLevel) bool {
	return fl.Field().IsZero()
}

func BooleanValidator(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.IsZero() {
		return true
	}
	if field.Kind() == reflect.Ptr && field.IsNil() {
		return true
	}
	if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Bool {
		return true
	}
	if field.Kind() == reflect.Bool {
		return true
	}
	return false
}

func DecimalGT(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.GreaterThan(baseValue)
}

func DecimalLT(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.LessThan(baseValue)
}
func DecimalGTE(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.GreaterThanOrEqual(baseValue)
}

func DecimalLTE(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.LessThanOrEqual(baseValue)
}

func PasswordValidator(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok {
		if strings.Contains(password, " ") {
			return false
		}

		if !uppercaseRegexp.MatchString(password) {
			return false
		}
		if !numberRegexp.MatchString(password) {
			return false
		}
		if !specialCharRegexp.MatchString(password) {
			return false
		}

		if len(password) < 8 || len(password) > 255 {
			return false
		}
		return true
	}

	return false
}

func CleanInputValidator(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().(string)
	if ok {
		if specialCharRegexp.MatchString(input) {
			return false
		}
		if len(input) < 4 || len(input) > 255 {
			return false
		}
		return true
	}
	return false
}

func PhoneNumberValidator(fl validator.FieldLevel) bool {
	phoneNumber, ok := fl.Field().Interface().(string)
	if ok {
		phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
		phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")

		if phoneRegexp.MatchString(phoneNumber) {
			return true
		}
	}

	return false
}

func TimeFormatValidator(fl validator.FieldLevel) bool {
	format := fl.Param()
	_, err := time.Parse(format, fl.Field().String())
	return err == nil
}

func DayOfWeekValidator(fl validator.FieldLevel) bool {
	day := fl.Field().String()
	return dayOfWeeks[day]
}

func NoDuplicatesValidator(fl validator.FieldLevel) bool {
	field := fl.Field()
	seen := make(map[interface{}]bool)

	for i := 0; i < field.Len(); i++ {
		value := field.Index(i).Interface()
		if seen[value] {
			return false
		}
		seen[value] = true
	}
	return true
}

func RoleValidator(fl validator.FieldLevel) bool {
	field, ok := fl.Field().Interface().(int64)
	if ok {
		if field == 1 || field == 3 {
			return true
		}
		return false
	}
	return false
}

func LatitudeValidator(fl validator.FieldLevel) bool {
	latitude, ok := fl.Field().Interface().(string)
	if ok {
		lat, err := decimal.NewFromString(latitude)
		if err != nil {
			return false
		}
		return lat.GreaterThanOrEqual(decimal.NewFromFloat(-90)) && lat.LessThanOrEqual(decimal.NewFromFloat(90))
	}
	return false
}

func LongitudeValidator(fl validator.FieldLevel) bool {
	longitude, ok := fl.Field().Interface().(string)
	if ok {
		long, err := decimal.NewFromString(longitude)
		if err != nil {
			return false
		}
		return long.GreaterThanOrEqual(decimal.NewFromFloat(-180)) && long.LessThanOrEqual(decimal.NewFromFloat(180))
	}
	return false
}

func Base64Validator(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if ok {
		_, err := decodeBase64(data)
		return err == nil
	}
	return false
}

// Decode Base64 (helper function for validation)
func decodeBase64(str string) ([]byte, error) {
	return []byte(str), nil
}
