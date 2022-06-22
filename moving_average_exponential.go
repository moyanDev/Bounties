
/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import "fmt"

/*
ExponentialMovingAverage returns the exponential average in _n_ number of periods

# Formula

![EMA formula](https://wikimedia.org/api/rest_v1/media/math/render/svg/05d06bdbee2c14031fd91ead6f5f772aec1ec964)

Where:

* _EMA<sub>t</sub>_ - is the value of the EMA at any time period _t_.
* _EMA<sub>t-1</sub>_ - is the value of the EMA at the previous period _t-1_.
* _p<sub>t</sub>_ - is the input value at a time period t.
* _α_ - is the coefficient that represents the degree of weighting decrease, a constant smoothing factor between 0 and 1.

_α_ is calculated with the following formula: