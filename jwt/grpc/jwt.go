// Package grpc ...
package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/lestrrat-go/jwx/jwt"
	gocorejwt "github.com/micro-business/go-core/jwt"
	"github.com/micro-business/go-core/system/errors"
)

// ParseAndVerifyToken parse the given token and verifies the signature
// ctx: Mandatory - the context
// jwksURL: Mandatory - the URL to where the public keys are stored
// verifyToken: Mandatory - indicates whether the provided token needs to be signature verified
// Returns either the parsed and verified token or error if something goes wrong
func ParseAndVerifyToken(ctx context.Context, jwksURL string, verifyToken bool) (jwt.Token, error) {
	if ctx == nil {
		return nil, errors.NewArgumentNilError("ctx", "ctx is required")
	}

	if verifyToken && len(jwksURL) == 0 {
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

	return gocorejwt.ParseAndVerifyToken(ctx, values[0], jwksURL, verifyToken)
}
