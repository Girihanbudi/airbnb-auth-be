package request

import "airbnb-auth-be/internal/pkg/validator"

type Me struct {
	UserId string `json:"userId" validation:"required"`
}

func (req *Me) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
