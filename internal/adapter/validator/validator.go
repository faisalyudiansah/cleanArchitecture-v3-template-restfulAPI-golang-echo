package validator

import (
	"fmt"

	model "server/internal/adapter/validator/model"
	"server/internal/adapter/validator/tools"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func (v *CustomValidator) ParseValidationErrors(err error) error {
	var errs model.ValidationErrors

	if err == nil {
		return errs
	}
	for _, err := range err.(validator.ValidationErrors) {
		e := model.ValidationError{
			Namespace:       err.Namespace(),
			Field:           err.Field(),
			StructNamespace: err.StructNamespace(),
			StructField:     err.StructField(),
			Tag:             err.Tag(),
			ActualTag:       err.ActualTag(),
			Kind:            fmt.Sprintf("%v", err.Kind()),
			Type:            fmt.Sprintf("%v", err.Type()),
			Value:           fmt.Sprintf("%v", err.Value()),
			Param:           err.Param(),
			Message:         err.Error(),
		}

		errs.Errors = append(errs.Errors, e)
	}

	return errs
}

func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterTagNameFunc(tools.TagNameFormatter)
	registerCommonRules(v)
	return &CustomValidator{
		Validator: v,
	}
}

func registerCommonRules(v *validator.Validate) {
	v.RegisterValidation("lt_today", dateOlderThanToday)
	v.RegisterValidation("lte_today", dateOlderOrEqualThanToday)
	v.RegisterValidation("gt_today", dateGreaterThanToday)
	v.RegisterValidation("gte_today", dateGreaterOrEqualThanToday)
	v.RegisterValidation("excluded", tools.ValidateExcluded)
	v.RegisterValidation("boolean", tools.BooleanValidator)
	v.RegisterValidation("dgte", tools.DecimalGT)
	v.RegisterValidation("dgte", tools.DecimalGTE)
	v.RegisterValidation("dlt", tools.DecimalLT)
	v.RegisterValidation("dlte", tools.DecimalLTE)
	v.RegisterValidation("password", tools.PasswordValidator)
	v.RegisterValidation("clean_input", tools.CleanInputValidator)
	v.RegisterValidation("phone_number", tools.PhoneNumberValidator)
	v.RegisterValidation("time_format", tools.TimeFormatValidator)
	v.RegisterValidation("day_of_weeks", tools.DayOfWeekValidator)
	v.RegisterValidation("no_duplicates", tools.NoDuplicatesValidator)
	v.RegisterValidation("role", tools.RoleValidator)
	v.RegisterValidation("latitude", tools.LatitudeValidator)
	v.RegisterValidation("longitude", tools.LongitudeValidator)
	v.RegisterValidation("base64", tools.Base64Validator)
}
