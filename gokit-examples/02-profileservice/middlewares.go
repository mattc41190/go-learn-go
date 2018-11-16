package profileservice

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware is a user defined type whose value must be a function which accepts an implementer of Service and returns an implementer of Service.
// Basically we are following a decorator pattern wherein the implementer of Service that was passed in will get decorated with additional fields
// or methods. The operation in practice will appear as if we are reassigning a variable over and over again to the result of different middlewares.
// By doing this a developer can increase added complexity of things like circuit breakers, loggers, metric collectors,
// and other generic service decorators without overloading the initial service.
type Middleware func(Service) Service

// loggingMiddleware is a struct which "wraps" or contains a service
// held on the `next` field. In this case it also maintains a field whose value is a `Logger`
// the `logger` field will be referenced and used to univarsally log things on the same value.
// loggingMiddleware will BECOME the service the application ends up using.
type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// PostProfile is a func which implements service.PostProfile by calling through the receiver's (mw)
// next field's PostProfile function and accomplishes some post processing via a defer func which
// uses loggingMiddleware's referred to in the func as its receiver alias `mw`
func (mw *loggingMiddleware) PostProfile(ctx context.Context, p Profile) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostProfile", "id", p.ID, "took", time.Since(begin), "err", err)
		// NOTE: I hate how err is implicitly set, smack of cleverness
	}(time.Now())
	return mw.next.PostProfile(ctx, p)
}

func (mw *loggingMiddleware) GetProfile(ctx context.Context, profileID string) (p Profile, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetProfile", "id", profileID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return mw.next.GetProfile(ctx, profileID)
}

func (mw *loggingMiddleware) PutProfile(ctx context.Context, profileID string, p Profile) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PutProfile", "id", profileID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PutProfile(ctx, profileID, p)
}

func (mw *loggingMiddleware) PatchProfile(ctx context.Context, profileID string, p Profile) (err error) {

	defer func(begin time.Time) {
		mw.logger.Log("method", "PatchProfile", "id", profileID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return mw.next.PatchProfile(ctx, profileID, p)
}

func (mw *loggingMiddleware) DeleteProfile(ctx context.Context, profileID string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteProfile", "id", profileID, "took", time.Since(begin), "err", err)
	}(time.Now())

	return mw.next.DeleteProfile(ctx, profileID)
}

// LoggingMiddleware accepts a Logger and returns a Middleware
// That is a PACKED statement. It takes an entity whose job is intelligent logging and returns a function
// This function as described above accepts an implementer of Service and returns an implementors of Service
// In the meantime the Service should be decorated with additional generic functionality that all Service methods should be able to
// closely simulate.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}
