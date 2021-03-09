/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"fmt"
)

/*
Mean returns the arithmetic mean value of the last n values.

# Formula

![SD formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/97821b8c43e3182faa22db06932846d1550866fb_

Where:

* _n_ - number of probes in observation.

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type Mean struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int
	count int

	sum float64

	// slice of data needed for calculation
	data []float64
}

// NewMean creates a new Mean with the given number of periods
// Example: NewMean(9)
func NewMean(n int) (*Mean, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &Mean{
		n: n,

		index: 0,
		count: 0,

		sum: 0,

		data: make([]float64, n),
	}, nil
}

// Next takes the next input and returns the next Mean valu