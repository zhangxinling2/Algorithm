package FixedWindowLimiter

import (
	"context"
	_ "embed"
	"errors"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"time"
)

//go:embed FixedWindowLimiter.lua
var FixedWindowLimiterLua string

type FixedWindowLimiter struct {
	limiter int
	ttl     int
	client  *redis.Client
	script  *redis.Script
}

func NewFixedWindowLimiter(client *redis.Client, limiter int, ttl time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limiter: limiter,
		ttl:     int(ttl / time.Millisecond),
		client:  client,
		script:  redis.NewScript(FixedWindowLimiterLua),
	}
}

func (l *FixedWindowLimiter) TryAcquire(ctx context.Context, resource string) error {
	success, err := l.script.Run(ctx, l.client, []string{resource}, l.limiter, l.ttl).Bool()
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Acquire Fail")
	}
	return nil
}
