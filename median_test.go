/*
Copyright 2020 Binh Nguyen
Licensed under terms of MIT license (see LICENSE)
*/

package tago

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestNewMedian(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *Median
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &Median{n: 9, data: make([]float64, 0, 9)}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewMedian(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.wa