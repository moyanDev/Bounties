
/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"fmt"
	"math"
)

/*
StandardDeviation returns the standard deviation of the last n values.

# Formula

![SD formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/2845de27edc898d2a2a4320eda5f57e0dac6f650)

Where:

* _σ_ - value of standard deviation for N given probes.
* _N_ - number of probes in observation.
* _x<sub>i</sub>_ - i-th observed value from N elements observation.

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
sd, _ := NewStandardDeviation(4)
sd.Next(10.)
sd.Next(20.)
sd.Next(30.)
sd.Next(20.)
sd.Next(10.)
```
*/
type StandardDeviation struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int
	count int

	m  float64
	m2 float64

	// slice of data needed for calculation
	data []float64
}

// NewStandardDeviation creates a new StandardDeviation with the given number of periods
// Example: NewStandardDeviation(9)
func NewStandardDeviation(n int) (*StandardDeviation, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters