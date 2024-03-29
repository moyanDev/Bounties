
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

func TestNewMinimum(t *testing.T) {
	data := make([]float64, 9)
	for i := range data {
		data[i] = math.Inf(1)
	}

	tests := map[string]struct {
		input   int
		want    *Minimum
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters},
		"zero n":     {input: 0, want: nil, wantErr: ErrInvalidParameters},
		"positive n": {input: 9, want: &Minimum{n: 9, data: data}, wantErr: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotSD, gotErr := NewMinimum(tc.input)
			if tc.wantErr != nil { // only check error returned if expecting one
				assert.EqualError(t, gotErr, tc.wantErr.Error(), "must return the correct error")
			}
			assert.Equal(t, tc.want, gotSD, "must return the correct value")
		})
	}
}

func TestMinimumNext(t *testing.T) {
	sd, _ := NewMinimum(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 4., want: 4.},
		{input: 1.2, want: 1.2},
		{input: 5., want: 1.2},
		{input: 3., want: 1.2},
		{input: 4., want: 3.},
		{input: 6., want: 3.},
		{input: 7., want: 4.},
		{input: 8., want: 6.},
		{input: -9., want: -9.},
		{input: 0., want: -9.},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := sd.Next(tc.input)
			diff := cmp.Diff(tc.want, got, floatComparer)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestMinimumReset(t *testing.T) {
	sd, _ := NewMinimum(3)
	tests := []struct {
		input float64
		want  float64
	}{
		{input: 5., want: 5.},
		{input: 7., want: 5.},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := sd.Next(tc.input)
			diff := cmp.Diff(tc.want, got, floatComparer)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}

	sd.Reset()
	diff := cmp.Diff(8., sd.Next(8.), floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}

func TestMinimumString(t *testing.T) {
	sd, _ := NewMinimum(4)
	want := "Min(4)"
	got := sd.String()
	diff := cmp.Diff(want, got, floatComparer)
	if diff != "" {
		t.Fatalf(diff)
	}
}