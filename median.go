/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
Median returns the median value of the last n values.

Where:

* _n_ - number of probes in observation.

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type Median struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	index int

	// slice of data needed for calculation
	data []float64
}

// NewMedian creates a new Median with the given number of periods
// Example: NewMedian(9)
func NewMedian(n int) (*Median, error) {
	if n <= 0 {
		return nil, ErrInvalidParameters
	}

	return &Median{
		n: n,

		index: 0,

		data: make([]float64, 0, n),
	}, nil
}

// Next takes the next input and returns the next Median value
func (m *Median) Next(input float64) float64 {
	// add input to data
	if len(m.data) < m.n {
		m.data = append(m.data, input)
	} else {
		m.index = m.index % m.n
		m.data[m.index] = input
	}
	m.index++
	return quickselectMedian(m.data, defaultPivotFunc)
}

// Reset reset