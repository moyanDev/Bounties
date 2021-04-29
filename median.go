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
	// number of periods (must be an integer greater t