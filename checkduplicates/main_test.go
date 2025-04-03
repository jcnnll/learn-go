package main

import "testing"

func Test_checkDuplicates(t *testing.T) {
	checkDuplicatesTests := []struct {
		name     string
		args     []int
		expected bool
	}{
		{"zero case", []int{}, false},
		{"single item", []int{5}, false},
		{"no duplicates", []int{1, 2, 3}, false},
		{"duplicates", []int{1, 1, 1, 2, 2}, true},
	}

	for _, e := range checkDuplicatesTests {
		result := checkDuplicates(e.args...)
		if result != e.expected {
			t.Errorf("%s: \texpected: %v\tgot:%v\n", e.name, e.expected, result)
		}
	}

}
