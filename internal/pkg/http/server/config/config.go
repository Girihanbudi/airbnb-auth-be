package config

type Config struct {
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	PrivateKey string `mapstructure:"privatekey"`
	PublicCert string `mapstructure:"publiccert"`
}
