package backoff

import (
	"time"
)

// RandFloat64 defines the methodology for jitter.
type RandFloat64 interface {
	// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
	Float64() float64
}

// Jitter implements the randomize backoff delays so that if a cluster of requests start at the same time, they won't operate in lockstep.
type Jitter struct {
	// Jitter is the randomized backoff of jitter
	Jitter time.Duration
	// Rand is make random numbers for jitter
	Rand RandFloat64
}

// BackOff returns the amount of time to wait before the next retry given the number of retries.
func (j *Jitter) BackOff(retries uint) time.Duration {
	backoff := float64(j.Jitter) * (j.Rand.Float64()*2 - 1)
	return time.Duration(backoff)
}
