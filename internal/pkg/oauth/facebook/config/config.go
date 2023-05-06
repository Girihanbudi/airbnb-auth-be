package config

type Config struct {
	ClientId     string   `mapstructure:"clientid"`
	ClientSecret string   `mapstructure:"clientsecret"`
	UserInfoApi  string   `mapstructure:"userinfoapi"`
	RedirectUrl  string   `mapstructure:"redirecturl"`
	Scopes       []string `mapstructure:"scopes"`
}
