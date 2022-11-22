package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tcs := []struct {
		desc string
		in1  int
		in2  int
		out  int
	}{
		{"positive", 1, 1, 2},
		{"negative", -1, -1, -2},
		{"positivenegative", -1, 3, 2},
	}

	for _, tt := range tcs {
		t.Run(tt.desc, func(t *testing.T) {
			out := add(tt.in1, tt.in2)
			if out != tt.out {
				t.Errorf("there is some error")
			}
		})
	}
}
