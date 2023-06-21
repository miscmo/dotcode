package limiter

import (
	"fmt"
	"testing"
	"time"
)

func TestSlidingWindowLimiter(t *testing.T)  {
	limiter := &SlidingWindowLimiter{}
	//limiter.Set(10, time.Second)

	for i := 0; i < 300 ; i++ {
		if limiter.Take() {
			fmt.Printf("task %v, time: %v\n", i, time.Now().Unix())
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
}