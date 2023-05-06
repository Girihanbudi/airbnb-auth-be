package request

import "airbnb-auth-be/internal/pkg/validator"

type RefreshToken struct {
	Token string `json:"token" validation:"required"`
}

func (req *RefreshToken) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
