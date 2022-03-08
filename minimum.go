
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
Minimum returns the lowest value in a given time frame

# Parameters

* _n_ - size of time frame (integer greater than 0)

# Example
```
```
*/
type Minimum struct {
	// number of periods (must be an integer greater than 0)
	n int

	// internal parameters for calculations
	minIndex int
	curIndex int

	// slice of data needed for calculation
	data []float64
}
