package main

import (
	"testing"
)

func TestAllIdentical(t *testing.T) {
	identicalArray := []int{1, 1, 1, 1, 1, 1, 1, 1}

	tests := []struct {
		testingArray []int
		expected     int
	}{
		{identicalArray[:], 1}, //array is all same and so only 1 should be returned
	}

	for _, test := range tests {
		result := removeDupes(test.testingArray)
		if result != test.expected {
			t.Errorf("For input %q, expected %d but got %d", test.testingArray, test.expected, result)
		}
	}
}
