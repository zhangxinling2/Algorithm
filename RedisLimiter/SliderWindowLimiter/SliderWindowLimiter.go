package SliderWindowLimiter

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//go:embed SliderWindowLimiter.lua
var SliderWindowLimiterLua string

type SliderWindowLimiter struct {
	ttl     int64
	limiter int
	client  *redis.Client
	script  *redis.Script
}

func NewSliderWindowLimiter(client *redis.Client, ttl time.Duration, limiter int) *SliderWindowLimiter {
	return &SliderWindowLimiter{
		ttl:     ttl.Milliseconds(),
		limiter: limiter,
		client:  client,
		script:  redis.NewScript(SliderWindowLimiterLua),
	}
}

func (l *SliderWindowLimiter) TryAcquire(ctx context.Context, resource string) error {
	now := time.Now().UnixNano()
	success, err := l.script.Run(ctx, l.client, []string{resource}, now-l.ttl, l.limiter, now, fmt.Sprintf("%d", now)).Bool()
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Acquire Fail")
	}
	return nil
}
