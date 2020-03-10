package backoff

import (
	"time"
)

// Linear implements the defined linear compensation algorithm.
type Linear struct {
	// BaseDelay is the amount of time to backoff after.
	BaseDelay time.Duration
	// Multiplier is the factor with which to multiply backoffs after a failed retry.
	Multiplier float64
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

// BackOff returns the amount of time to wait before the next retry given the number of retries.
func (l *Linear) BackOff(retries uint) time.Duration {
	backoff := l.BaseDelay + time.Duration(float64(l.BaseDelay)*l.Multiplier*float64(retries))
	if backoff > l.MaxDelay {
		backoff = l.MaxDelay
	}
	return backoff
}
