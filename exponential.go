package backoff

import (
	"math"
	"time"
)

// Exponential implements the defined exponential compensation algorithm.
type Exponential struct {
	// BaseDelay is the amount of time to backoff after the first failure.
	BaseDelay time.Duration
	// Multiplier is the factor that exponentially multiplies the backoff after a retry failure.
	Multiplier float64
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

// BackOff returns the amount of time to wait before the next retry given the number of retries.
func (e *Exponential) BackOff(retries uint) time.Duration {
	backoff := time.Duration(float64(e.BaseDelay) * math.Pow(e.Multiplier, float64(retries)))
	if backoff > e.MaxDelay {
		backoff = e.MaxDelay
	}
	return backoff
}
