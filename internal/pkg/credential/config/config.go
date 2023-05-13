package config

type Config struct {
	PrivateKey string `mapstructure:"privatekey"`
	PublicCert string `mapstructure:"publiccert"`
}
