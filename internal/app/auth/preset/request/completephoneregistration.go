package request

import (
	"airbnb-auth-be/internal/pkg/validator"
	"time"
)

type CompletePhoneRegistration struct {
	Otp         string `json:"otp" validation:"required,numeric,len=6"`
	FirstName   string `json:"firstName" validation:"required"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" validation:"required,email"`
	DateOfBirth string `json:"dateOfBirth" validation:"required,date_iso8601"`
}

func (req *CompletePhoneRegistration) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (req *CompletePhoneRegistration) ConvertedDateOfBirth() time.Time {
	dob, _ := time.Parse("2006-01-02", req.DateOfBirth)
	return dob
}
