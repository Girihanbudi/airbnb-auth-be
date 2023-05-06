package request

import (
	"airbnb-auth-be/internal/pkg/validator"
)

type MakePhoneSession struct {
	Otp string `json:"otp" validation:"required,numeric,len=6"`
}

func (req *MakePhoneSession) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
