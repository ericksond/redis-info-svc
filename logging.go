package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   RedisInfoService
}

func (mw loggingMiddleware) Info(addr string, passwd string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "info",
			"input", addr,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Info(addr, passwd)
	return
}
