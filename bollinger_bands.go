/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

/*
BollingerBands
The Bollinger Bands are represented by Average EMA and standard deviaton that is moved 'k' times away in both directions from calculated average value.

# Formula

See SMA, SD documentation.

* _BB<sub>Middle Band</sub>_ - Simple Moving Average (SMA).
* _BB<sub>Upper Band</sub>_ = SMA + SD of observation * multipler (usually 2.0)
* _BB<sub>Lower Band</sub>_ = SMA - SD of observation * multipler (usually 2.0)

# Parameters

* _n_ - number of periods (integer greater than 0)

# Example
```
```
*/
type BollingerBands struct {
	// number of periods (must be an integer greater than 0)
	n