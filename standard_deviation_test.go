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

func TestNewStandardDeviation(t *testing.T) {
	tests := map[string]struct {
		input   int
		want    *StandardDeviation
		wantErr error
	}{
		"negative n": {input: -3, want: nil, wantErr: ErrInvalidParameters