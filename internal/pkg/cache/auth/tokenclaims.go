package auth

import "time"

type DefaultClaims struct {
	UserID      string     `json:"userId"`
	FirstName   string     `json:"first_name,omitempty"`
	FullName    string     `json:"full_name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	CountryCode *int       `json:"country_code,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Image       string     `json:"image,omitempty"`
	Password    *string    `json:"password,omitempty"`
	Role        string     `json:"role,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`

	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	VerifiedAt *time.Time `json:"verified_at,omitempty"`
}
