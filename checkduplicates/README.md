# Check for duplicate args problem

Implement a variadic function called checkDuplicates that accepts any number of integer values and checks for duplication.

## Concrete examples
- `checkDuplicates()` has 0 duplicate elements and returns `false`
- `checkDuplicates(1, 2, 3)` has 0 duplicates and returns `false`
- `checkDuplicates(1, 2, 1)` has 1 duplicate and returns `true`

## Rules
- If there are no args then there are no duplicates
- The input args can be in any order
- If there are any duplicate args return true

The following test cases will determine whether the above rules are applied:

```
		{"zero case", []int{}, false},
		{"single item", []int{5}, false},
		{"no duplicates", []int{1, 2, 3}, false},
		{"duplicates", []int{1, 1, 1, 2, 2}, true},
```

## Solution
- To achieve time complexity of O(n) no nested loop is used to compare elements in the array
- Because order is random, use a frequency counter to track duplication
- Short circuit the function by returning false if there is 1 or 0 args
- Loop through the args once and add the frequency count for each arg
- While creating the frequecncy counter, optimise time by checking if the count is 2 for the current count value and return true
```go
func checkDuplicates(xi ...int) bool {
	if len(xi) < 2 {
		return false
	}
	fc := make(map[int]int)

	for _, v := range xi {
		if _, ok := fc[v]; ok {
			fc[v] += 1
			if fc[v] == 2 {
				return true
			}
			continue
		}
		fc[v] = 1
	}
	return false
}

```
