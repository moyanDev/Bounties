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
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &Maximum{n: 9, data: data}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewMaximum(tc.input)
			if tc.wantErr != nil { // 