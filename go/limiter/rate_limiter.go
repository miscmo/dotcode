package limiter


import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate int
	begin time.Time
	cycle time.Duration
	count int
	lock sync.Mutex
}


func (limiter *RateLimiter) Take() bool {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()

	if limiter.count <= limiter.rate {
		limiter.count++
		return true
	}

	now := time.Now()
	if now.Sub(limiter.begin) > limiter.cycle {
		limiter.Reset(now)
		return true
	}

	return false
}

func (limiter *RateLimiter) Set(rate int, cycle time.Duration) {
	limiter.rate = rate
	limiter.cycle = cycle
	limiter.Reset(time.Now())
}

func (limiter *RateLimiter) Reset(begin time.Time) {
	limiter.begin = begin
	limiter.count = 0
}
