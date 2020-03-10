package backoff

import (
	"testing"
	"time"
)

var testLinear = &Linear{
	BaseDelay:  1 * time.Second,
	Multiplier: 2,
	MaxDelay:   30 * time.Second,
}

var testsLinear = []struct {
	retries uint
	want    time.Duration
}{
	{0, testLinear.BaseDelay * 1},
	{1, testLinear.BaseDelay * 3},
	{2, testLinear.BaseDelay * 5},
	{3, testLinear.BaseDelay * 7},
	{14, testLinear.BaseDelay * 29},
	{15, testLinear.BaseDelay * 30},
	{16, testLinear.BaseDelay * 30},
}

func TestLinear_Backoff(t *testing.T) {
	for _, tt := range testsLinear {
		t.Run("", func(t *testing.T) {
			if got := testLinear.BackOff(tt.retries); got != tt.want {
				t.Errorf("DefaultExponentialStrategy.BackOff() = %v, want %v", got, tt.want)
			}
		})
	}
}
