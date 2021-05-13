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
func NewMedian(n i