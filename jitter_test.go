package backoff

import (
	"testing"
	"time"
)

func TestJitter(t *testing.T) {
	j := &Jitter{time.Second, 0.2, grand}
	for i := 0; i != 10; i++ {
		t.Run("", func(t *testing.T) {
			want := float64(j.Amplitude)
			wantMin := time.Duration(want * (1 - j.Randomized))
			wantMax := time.Duration(want * (1 + j.Randomized))
			for i := 0; i != 10; i++ {
				if got := j.BackOff(uint(i)); got < wantMin || got > wantMax {
					t.Errorf("DefaultExponentialStrategy.BackOff() = %v, want between %v, %v", got, wantMin, wantMax)
					break
				}
			}
		})
	}
}
