package appcontext

const minute = 60
const hour = minute * 60
const day = hour * 24

const (
	LocaleCode     = "locale" // locale application context code
	LocaleDuration = 7 * day  // expires in a week
	LocaleDefault  = "en"     // default language using English US
)

const (
	CurrencyCode     = "currency" // currency application context code
	CurrencyDuration = 7 * day    // expires in a week
	CurrencyDefault  = "USD"      // default currency using United State Dollar
)

const (
	IsLoggedInCode       = "logged_in"   // user logged in status
	AccessTokenCode      = "__Secure.at" // access token application context code
	AccessTokenDuration  = 15 * minute   // expires in 15 minutes
	RefreshTokenCode     = "__Secure.rt" // refresh token application context code
	RefreshTokenDuration = 7 * day       // expires in a week
	OauthCode            = "oauth"       // oauth application context code
	OauthDuration        = 1 * minute    // expires in a minute
	OtpDuration          = 1 * minute    // expires in a minute
)

const (
	UserClaims = "user_claims" // user id application context code
)
