package response

import "time"

type UserDefaultSetting struct {
	Locale   string `json:"locale"`
	Currency string `json:"currency"`
}

type Account struct {
	Provider     string    `json:"provider"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiredAt    time.Time `json:"expiredAt"`
	TokenType    string    `json:"tokenType"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Me struct {
	FirstName   string    `json:"firstName"`
	FullName    string    `json:"fullName"`
	Email       *string   `json:"email"`
	CountryCode *int      `json:"countryCode"`
	PhoneNumber *string   `json:"phoneNumber"`
	Image       string    `json:"image"`
	Role        string    `json:"role"`
	DateOfBirth time.Time `json:"dateOfBirth"`

	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	VerifiedAt *time.Time `json:"verifiedAt"`

	DefaultSetting *UserDefaultSetting `json:"defaultSetting"`
	Accounts       *[]Account          `json:"accounts"`
}
