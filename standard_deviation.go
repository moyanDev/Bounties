
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