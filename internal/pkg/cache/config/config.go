package config

type CacheDefaultConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type Config struct {
	Auth CacheDefaultConfig `mapstructure:"auth"`
	Otp  CacheDefaultConfig `mapstructure:"otp"`
}
