package main

import "testing"

func Test_maxSubarraySum(t *testing.T) {
	maxSubarraySumTests := []struct {
		name        string
		testSlice   []int
		testWinSize int
		expected    int
	}{
		{"zero case", []int{}, 4, 0},
		{"windown length 2 array length 7", []int{1, 2, 5, 2, 8, 1, 5}, 2, 10},
		{"window length 4 array length 7", []int{1, 2, 5, 2, 8, 1, 5}, 4, 17},
		{"window length 1 array length 4", []int{4, 2, 1, 6}, 1, 6},
		{"window length 4 array length 5", []int{4, 2, 1, 6, 2}, 4, 13},
	}

	for _, e := range maxSubarraySumTests {
		result := maxSubarraySum(e.testSlice, e.testWinSize)
		if result != e.expected {
			t.Errorf("%s: \t expected: %d\tgot:%d\n", e.name, e.expected, result)
		}
	}
}
