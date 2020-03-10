package backoff

import (
	"time"
)

type Superposition []Strategy

// BackOff returns the amount of time to wait before the next retry given the number of retries.
func (s *Superposition) BackOff(retries uint) time.Duration {
	var backoff time.Duration
	for _, strategy := range *s {
		backoff += strategy.BackOff(retries)
	}
	return backoff
}
