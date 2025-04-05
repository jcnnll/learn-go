# Count unique elements problem

Implement a function called countUnique that accepts a sorted array and counts the unique occurrances for each elemet in the array.
There can be negative numbers in the array, but it will always be sorted.

## Concrete examples
- The array `[]` has 0 unique elements
- The array `[1]` has 1 unique element
- The array `[1,1,1,1,1,2]` has 2 unique elements
- The array `[1,2,3,4,4,4,7,7,12,12,13]` has 7 unique elements
- The array `[-2,-1,-1,0,1]` has 4 unique elements

## Rules
- If the array is empty return 0
- The input array will always be sorted

The following test cases will determine whether the above rules are applied:

```
		{"zero case", []int{}, 0},
		{"single item", []int{5}, 1},
		{"2 items only", []int{1, 2}, 2},
		{"2 items with repeated elements", []int{1, 1, 1, 2, 2}, 2},
		{"expecting count of 7 elements", []int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}, 7},
		{"expecting count of 4", []int{-2, -1, -1, 0, 1}, 4},
```

## Solution
- Check the length of the array, if it is empty, return 0
- If there is only 1 element in the array, then 1 unique element exists
- Use 2 pointers to track and evaluate each element, move unique elements to the first pointer position and overwrite duplicate elements
- To keep space complexity to O(1) use the input array to store the elements in the input array
- To achieve time complexity of O(n) no nested loop is used to compare elements in the array
```go

func countUnique(xi []int) int {
	if len(xi) == 1 {
		return 1
	}
	if len(xi) == 0 {
		return 0
	}
	i := 0
	for j := range xi[0:] {
		if xi[i] != xi[j] {
			i++
			xi[i] = xi[j]
		}
	}
	return i + 1
}
```
