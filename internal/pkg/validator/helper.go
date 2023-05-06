package validator

func ValidateStruct(data interface{}) (validatorErr error) {
	validate := GetValidator()
	err := validate.Struct(data)

	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	// if err != nil {
	// // get translation
	// if lang == nil {
	// 	defaultLang := locale.GetDefaultLang()
	// 	lang = &defaultLang
	// }
	// localizer := locale.NewLocalizer(*lang)
	// msg := i18n.LocalizeConfig{MessageID: "validation_default"}
	// t, err := localizer.Localize(&msg)

	// errs := err.(goValidator.ValidationErrors)
	// // create list of validation message
	// for _, err := range errs {
	// 	errMessages = append(errMessages, fmt.Sprintf(t, err.Field()))
	// }
	// }

	return err
}
