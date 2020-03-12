package backoff

import (
	"testing"
	"time"
)

func TestJitter(t *testing.T) {
	j := &Jitter{time.Second, grand}
	for i := 0; i != 10; i++ {
		t.Run("", func(t *testing.T) {
			wantMin := -j.Jitter
			wantMax := j.Jitter
			for i := 0; i != 10; i++ {
				if got := j.BackOff(uint(i)); got < wantMin || got > wantMax {
					t.Errorf("DefaultExponentialStrategy.BackOff() = %v, want between %v, %v", got, wantMin, wantMax)
					break
				}
			}
		})
	}
}
