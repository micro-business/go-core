// Package fasthttp ...
package fasthttp

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lestrrat-go/jwx/jwt"
	gocorejwt "github.com/micro-business/go-core/jwt"
	"github.com/micro-business/go-core/system/errors"
	fasthttppkg "github.com/valyala/fasthttp"
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

	convertedCtx, ok := ctx.(*fasthttppkg.RequestCtx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to cast ctx to fasthttp.RequestCtx")
	}

	return gocorejwt.ParseAndVerifyToken(ctx, string(convertedCtx.Request.Header.Peek(fasthttppkg.HeaderAuthorization)), jwksURL, verifyToken)

}
