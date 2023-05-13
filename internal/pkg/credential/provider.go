package credential

import (
	"airbnb-auth-be/internal/pkg/credential/config"
	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

type Options struct {
	config.Config
}

type TlsCredentials struct {
	Options
	Tls *credentials.TransportCredentials
}

func NewTLSCredentials(options Options) (creds TlsCredentials) {
	creds.Options = options

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(options.PublicCert, options.PrivateKey)
	if err != nil {
		return creds
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	tls := credentials.NewTLS(config)
	creds.Tls = &tls

	return creds
}
