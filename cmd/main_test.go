package main_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	tcs := []struct {
		name string
	}{
		{"a"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Print("what")
			fmt.Print("testing")
		})
	}
}
