// Package util implements different utilities
package util

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/micro-business/go-core/system/errors"
)

const bearerTokenPrefix = "Bearer "

// ParseAndVerifyToken parse the given token and verifies the signature
// ctx: Mandatory - the context
// jwksURL: Mandatory - the URL to where the public keys are stored
// Returns either the parsed and verified token or error if something goes wrong
func ParseAndVerifyToken(ctx context.Context, jwksURL string) (jwt.Token, error) {
	if ctx == nil {
		return nil, errors.NewArgumentNilError("ctx", "ctx is required")
	}

	if len(jwksURL) == 0 {
		return nil, errors.NewArgumentNilError("jwksURL", "jwksURL is required")
	}

	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := metadata["authorization"]
	if len(values) == 0 || len(values[0]) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	authorizationToken := values[0]
	if !strings.HasPrefix(authorizationToken, bearerTokenPrefix) {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token format is not Bearer")
	}

	bearerToken := authorizationToken[len(bearerTokenPrefix):]
	if len(bearerToken) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

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
