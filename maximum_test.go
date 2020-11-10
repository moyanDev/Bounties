/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestNewMaximum(t *testing.T) {
	data := make([]float64, 9)
	for i := range data {
		data[i] = math.Inf(-1)
	}

	tests := map[string]struct {
		input   int
		want    *Maximum
		wantErr error
	}{
		"