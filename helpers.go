/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"math"

	"github.com/google/go-cmp/cmp"
)

const (
	tolerance = 0.001
)

var floatComparer = cmp.Co