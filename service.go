package main

import (
	"errors"
	"regexp"
	"strings"

	redis "gopkg.in/redis.v4"
)

// RedisInfoService provides output from the redis-cli info command
type RedisInfoService interface {
	Info(string, string) (map[string]interface{}, error)
}

type redisInfoService struct{}

func (redisInfoService) Info(addr string, passwd string) (map[string]interface{}, error) {
	if addr == "" {
		return nil, ErrHostEmpty
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})

	info, err := client.Info().Result()

	m := make(map[string]interface{})
	x := strings.Split(info, "\r\n")

	for i := 0; i < len(x); i++ {
		match, err := regexp.MatchString(":", x[i])
		if err != nil {
			return nil, ErrRegexMatch
		}

		if match == true {
			y := strings.Split(x[i], ":")
			m[y[0]] = y[1]
		}
	}

	return m, err
}

var (
	ErrHostEmpty    = errors.New("empty host request")
	ErrRegexMatch   = errors.New("regex match error")
	ErrMarshalError = errors.New("marshal error")
)
