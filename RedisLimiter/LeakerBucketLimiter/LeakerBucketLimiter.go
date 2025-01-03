package LeakerBucketLimiter

import (
	"context"
	_ "embed"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

//go:embed LeakerBucketLImiter.lua
var LeakerBucketLimiterLua string

type LeakerBucketLimiter struct {
	peakLevel int
	speed     int
	client    *redis.Client
	script    *redis.Script
}

func NewLeakerBucketLimiter(client *redis.Client, peakLevel, speed int) *LeakerBucketLimiter {
	return &LeakerBucketLimiter{
		peakLevel: peakLevel,
		speed:     speed,
		client:    client,
		script:    redis.NewScript(LeakerBucketLimiterLua),
	}
}

func (l *LeakerBucketLimiter) TryAcquire(ctx context.Context, resource string) error {
	success, err := l.script.Run(ctx, l.client, []string{resource}, l.peakLevel, l.speed, time.Now().Second()).Bool()
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Acquire Fail")
	}
	return nil
}
