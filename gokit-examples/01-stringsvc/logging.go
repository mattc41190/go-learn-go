package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type middlewareLogging struct {
	logger log.Logger
	next   StringService
}

func (mw middlewareLogging) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Uppercase(s)
	return
}

func (mw middlewareLogging) Count(s string) (output int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.Count(s)
	return
}
