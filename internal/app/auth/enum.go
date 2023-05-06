package auth

// Provider is a unique name of which provider that supplies the OAuth service
type Provider int

const (
	ProviderEmail Provider = iota
	ProviderPhone
	ProviderGoogle
	ProviderFacebook
)

var keys = []string{"email", "phone", "google", "facebook"}
var providers = []Provider{ProviderEmail, ProviderPhone, ProviderGoogle, ProviderFacebook}

func (m Provider) String() string {
	return keys[m]
}

func ProviderFromKey(lookupKey string) *Provider {
	for i, key := range keys {
		if key == lookupKey {
			return &providers[i]
		}
	}

	return nil
}
