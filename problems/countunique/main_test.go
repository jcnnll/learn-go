package main

import (
	"testing"
)

func Test_countUnique(t *testing.T) {
	countUniqueTests := []struct {
		name      string
		testSlice []int
		expected  int
	}{
		{"zero case", []int{}, 0},
		{"single item", []int{5}, 1},
		{"2 items only", []int{1, 2}, 2},
		{"2 items with repeated elements", []int{1, 1, 1, 2, 2}, 2},
		{"expecting count of 7 elements", []int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}, 7},
		{"expecting count of 4", []int{-2, -1, -1, 0, 1}, 4},
	}

	for _, e := range countUniqueTests {
		result := countUnique(e.testSlice)
		if e.expected != result {
			t.Errorf("%s: expected %d but got %d", e.name, e.expected, result)
		}
	}

}
