package request

import "airbnb-auth-be/internal/pkg/validator"

type SignOut struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken" validation:"required"`
}

func (req *SignOut) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
