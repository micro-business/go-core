// Package jwt ...
package jwt

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/micro-business/go-core/system/errors"
)

const bearerTokenPrefix = "Bearer "

// ParseAndVerifyToken parse the given token and verifies the signature
// ctx: Mandatory - the context
// authorizationToken: Mandatory - the auhorization token
// jwksURL: Mandatory - the URL to where the public keys are stored
// verifyToken: Mandatory - indicates whether the provided token needs to be signature verified
// Returns either the parsed and verified token or error if something goes wrong
func ParseAndVerifyToken(ctx context.Context, authorizationToken string, jwksURL string, verifyToken bool) (jwt.Token, error) {
	if len(authorizationToken) == 0 {
		return nil, errors.NewArgumentNilError("authorizationToken", "authorizationToken is required")
	}

	if !strings.HasPrefix(authorizationToken, bearerTokenPrefix) {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token format is not Bearer")
	}

	bearerToken := authorizationToken[len(bearerTokenPrefix):]
	if len(bearerToken) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	if verifyToken {
		keySet, err := jwk.Fetch(ctx, jwksURL)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Fail to fetch the key set")
		}

		token, err := jwt.Parse([]byte(bearerToken), jwt.WithKeySet(keySet))
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Failed to parse and validate the received token")
		}

		return token, nil
	}

	token, err := jwt.Parse([]byte(bearerToken))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to parse the received token")
	}

	return token, nil
}
