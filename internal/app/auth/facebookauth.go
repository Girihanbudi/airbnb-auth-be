package auth

// User information provided by facebook apis
type FacebookUserInfo struct {
	Id        string                            `json:"id"`
	FirstName string                            `json:"first_name"`
	Name      string                            `json:"name"`
	Email     string                            `json:"email"`
	Picture   map[string]map[string]interface{} `json:"picture"`
	Gender    string                            `json:"gender"`
	Verified  bool                              `json:"verified"`
	Locale    string                            `json:"locale"`
}
