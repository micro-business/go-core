// Package middleware implements different go-kit middlewares
package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
)

type middlewareProviderService struct {
	logger            *zap.Logger
	logMessageDetails bool
	dateTimeFormat    string
}

// NewMiddlewareProviderService creates new instance of the middlewareProviderService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// logger: Indicates whether the request and response message details sholuld be logged
// dateTimeFormat: Optional. If provided, will be used to format the date time object
// Returns the new service or error if something goes wrong
func NewMiddlewareProviderService(
	logger *zap.Logger,
	logMessageDetails bool,
	dateTimeFormat string) (MiddlewareProviderContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if dateTimeFormat == "" {
		dateTimeFormat = time.RFC3339Nano
	}

	return &middlewareProviderService{
		logger:            logger,
		logMessageDetails: logMessageDetails,
		dateTimeFormat:    dateTimeFormat,
	}, nil
}

// CreateLoggingMiddleware creates the logging middleware.
// endpointName: Mandatory. The name of the endpoint
// Returns the new endpoint with logging middleware added
func (service *middlewareProviderService) CreateLoggingMiddleware(endpointName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				end := time.Now()
				duration := time.Since(begin)
				fields := []zap.Field{
					zap.String("endpoint_name", endpointName),
					zap.String("start_time", begin.Format(service.dateTimeFormat)),
					zap.String("end_time", end.Format(service.dateTimeFormat)),
					zap.String("duration", duration.String()),
				}

				if service.logMessageDetails {
					fields = append(fields, zap.Any("request", request), zap.Any("response", response))
				}

				if err == nil {
					service.logger.Info(
						"Endpoint call succeeded",
						fields...)
				} else {
					fields = append(fields, zap.Error(err))
					service.logger.Error(
						"Endpoint call failed",
						fields...)
				}
			}(time.Now())

			return next(ctx, request)
		}
	}
}
