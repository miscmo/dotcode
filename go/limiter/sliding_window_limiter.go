package limiter

import (
	"sync"
	"time"
)

type timeSlot struct {
	timestamp time.Time
	count int
}

func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}

	return count
}


type SlidingWindowLimiter struct {
	mu sync.Mutex
	SlotDuration time.Duration
	WinDuration time.Duration
	numSlots int
	windows []*timeSlot
	maxReq int
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration: winDuration,
		numSlots: int(winDuration / slotDuration),
		maxReq: maxReq,
	}
}

func (l *SlidingWindowLimiter) Take() bool {

	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	timeoutOffset := -1

	for i, ts := range l.windows {
		if ts.timestamp.Add(l.WinDuration).After(now) {

		}
	}

	return false
}
