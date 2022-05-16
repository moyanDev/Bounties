/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import "fmt"

/*
MovingAverage returns the average in _n_ number of periods

# Formula

![SMA](https://wikimedia.org/api/rest_v1/media/math/render/svg/e2bf09dc6deaf86b3607040585fac6078f9c7c89)

Where:

* _SMA<sub>t</sub>_ - value of simple moving average at a point of time _t_
* _n_ - number of periods (length)
* _p<sub>t</sub>_ - input value at a point of time _t_

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type MovingAverage struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int
	count int

	sum float64

	// slice of data needed for calculation
	data []float64
}

// NewMovingAverage creates a new MovingAverage with the given number of periods
// Example: NewMovingAverage(9)
func NewMovingAverage(n int) (*MovingAverage, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &MovingAverage{
		n: n,

		index: 0,
		count: 0,

		sum: 0,

		data: make