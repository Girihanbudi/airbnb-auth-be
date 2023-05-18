package config

type Config struct {
	Addresses []string `mapstructure:"addresses"`
	Username  string   `mapstructure:"username"`
	Password  string   `mapstructure:"password"`
	Separator string   `mapstructure:"separator"`
}
