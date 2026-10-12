package auth

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"employeejwt/internal/config"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type TokenValidator struct {
	validator *validator.Validator
}

func NewTokenValidator() (*TokenValidator, error) {
	issuer := config.Auth0Issuer()
	audience := config.Auth0Audience()
	if issuer == "" || audience == "" {
		return nil, fmt.Errorf("AUTH0_ISSUER and AUTH0_AUDIENCE must both be configured")
	}

	issuerURL, err := url.Parse(issuer)
	if err != nil {
		return nil, fmt.Errorf("invalid AUTH0_ISSUER: %w", err)
	}
	if issuerURL.Scheme != "https" || issuerURL.Host == "" {
		return nil, fmt.Errorf("AUTH0_ISSUER must be an HTTPS issuer URL")
	}

	keyProvider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	return newTokenValidator(issuer, audience, keyProvider.KeyFunc)
}

func newTokenValidator(issuer, audience string, keyFunc func(context.Context) (interface{}, error)) (*TokenValidator, error) {
	jwtValidator, err := validator.New(
		keyFunc,
		validator.RS256,
		issuer,
		[]string{audience},
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("configure Auth0 token validator: %w", err)
	}

	return &TokenValidator{validator: jwtValidator}, nil
}

func (v *TokenValidator) ValidateToken(tokenString string) error {
	_, err := v.validator.ValidateToken(context.Background(), tokenString)
	if err != nil {
		return fmt.Errorf("validate Auth0 access token: %w", err)
	}
	return nil
}
