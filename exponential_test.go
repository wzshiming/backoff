package backoff

import (
	"math"
	"testing"
	"time"
)

var testExponential = &Exponential{
	BaseDelay:  1 * time.Second,
	Multiplier: 1.6,
	MaxDelay:   120 * time.Second,
}

var testsExponential = []struct {
	retries uint
	want    float64
}{
	{0, math.Pow(testExponential.Multiplier, 0)},
	{1, math.Pow(testExponential.Multiplier, 1)},
	{2, math.Pow(testExponential.Multiplier, 2)},
	{3, math.Pow(testExponential.Multiplier, 3)},
	{4, math.Pow(testExponential.Multiplier, 4)},
	{5, math.Pow(testExponential.Multiplier, 5)},
	{6, math.Pow(testExponential.Multiplier, 6)},
	{7, math.Pow(testExponential.Multiplier, 7)},
	{8, math.Pow(testExponential.Multiplier, 8)},
	{9, math.Pow(testExponential.Multiplier, 9)},
	{10, math.Pow(testExponential.Multiplier, 10)},
	{11, 120},
	{12, 120},
	{13, 120},
}

func TestExponential_Backoff(t *testing.T) {
	for _, tt := range testsExponential {
		t.Run("", func(t *testing.T) {
			want := time.Duration(tt.want * float64(time.Second))
			if got := testExponential.BackOff(tt.retries); got != want {
				t.Errorf("DefaultExponentialStrategy.BackOff() = %v, want %v", got, want)
			}
		})
	}
}
