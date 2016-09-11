package main

import (
	"errors"

	redis "gopkg.in/redis.v4"
)

// RedisInfoService provides output from the redis-cli info command
type RedisInfoService interface {
	Info(string, string) (string, error)
}

type redisInfoService struct{}

func (redisInfoService) Info(addr string, passwd string) (string, error) {
	if addr == "" {
		return "", ErrHostEmpty
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       0,
	})

	info, err := client.Info().Result()
	return info, err
}

// ErrHostEmpty is returned when an input string is empty
var ErrHostEmpty = errors.New("empty host request")
