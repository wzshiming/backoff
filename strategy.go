package backoff

import (
	"math/rand"
	"time"
)

// Strategy defines the methodology for backing off after.
type Strategy interface {
	// BackOff returns the amount of time to wait before the next retry given the number of consecutive failures.
	BackOff(retries uint) time.Duration
}

var DefaultLinear Strategy = &Superposition{
	defaultJitter,
	&Linear{
		BaseDelay:  1 * time.Second,
		Multiplier: 2,
		MaxDelay:   120 * time.Second,
	},
}

var DefaultExponential Strategy = &Superposition{
	defaultJitter,
	&Exponential{
		BaseDelay:  1 * time.Second,
		Multiplier: 1.6,
		MaxDelay:   120 * time.Second,
	},
}

var DefaultConstant Strategy = &Superposition{
	defaultJitter,
	&Constant{
		Delay: 1 * time.Second,
	},
}

var defaultJitter Strategy = &Jitter{
	Amplitude:  1 * time.Second,
	Randomized: 0.2,
	Rand:       grand,
}

var grand = rand.New(rand.NewSource(time.Now().UnixNano()))
