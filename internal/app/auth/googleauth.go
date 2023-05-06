package auth

// User information provided by google apis
type GoogleUserInfo struct {
	Id            string `json:"id"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
	Locale        string `json:"locale"`
}
