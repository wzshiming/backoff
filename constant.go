package backoff

import (
	"time"
)

// Constant implements the defined constant compensation algorithm.
type Constant struct {
	// Delay is the amount of time to backoff after.
	Delay time.Duration
}

// BackOff returns the amount of time to wait before the next retry given the number of retries.
func (c *Constant) BackOff(retries uint) time.Duration {
	return c.Delay
}
