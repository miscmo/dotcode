package limiter

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T)  {
	limiter := &RateLimiter{}
	limiter.Set(10, time.Second)

	runtime.GOMAXPROCS(2)

	for i := 0; i < 300 ; i++ {
		if limiter.Take() {
			fmt.Printf("task %v, time: %v\n", i, time.Now().Unix())
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
}