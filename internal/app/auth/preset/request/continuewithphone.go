package request

import (
	"airbnb-auth-be/internal/pkg/validator"
	"errors"
)

type ContinueWithPhone struct {
	CountryCode int    `json:"countryCode" validation:"required"`
	PhoneNumber string `json:"phoneNumber" validation:"required,max=12,numeric"`
}

func (req *ContinueWithPhone) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	if req.PhoneNumber[0:1] == "0" {
		return false, errors.New("phone number with leading zero not allowed")
	}

	return true, nil
}
