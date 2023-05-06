package jwt

import (
	"airbnb-auth-be/internal/pkg/env"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateToken(duration int, additionalCreds *map[string]interface{}) (token string, claims jwt.MapClaims, err error) {
	claims = jwt.MapClaims{}
	tokenId, generateIdErr := gonanoid.New()
	if generateIdErr != nil {
		err = generateIdErr
		return
	}

	expiry := time.Now().Add(time.Duration(duration / 60 * int(time.Minute)))

	claims["jti"] = tokenId                    // token unique id
	claims["iss"] = env.CONFIG.Domain          // issuer
	claims["exp"] = jwt.NewNumericDate(expiry) // expired time

	// added additional claims
	if additionalCreds != nil {
		for k, v := range *additionalCreds {
			if _, ok := claims[k]; ok {
				claims[k] = v
			}
		}
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the generated key using secretKey
	key := []byte(env.CONFIG.Jwt.Secret)
	token, signWithKeyErr := jwtToken.SignedString(key)
	if signWithKeyErr != nil {
		err = signWithKeyErr
		return
	}

	return
}

func ExtractTokenMetadata(token string) *jwt.MapClaims {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unknown signing method")
		}
		key := []byte(env.CONFIG.Jwt.Secret)
		return key, nil
	})
	if err != nil || !jwtToken.Valid {
		return nil
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	return &claims
}
