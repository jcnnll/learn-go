# Max sum sub array problem

Write a function named maxSubarraySum that accepts an array of integers as an input parameter and a number of type integer (n).
The function will calculate and return the maximum sum of n consecutive elements of the array

## Concrete examples
- The array `[1, 2, 5, 2, 8, 1, 5]` where n is `2` should return 10
- The array `[1, 2, 5, 2, 8, 1, 5]` where n is `4` should return 17
- The array `[4, 2, 1, 6]` where n is `1` should return 6
- The array `[4, 2, 1, 6, 2]` where n is `4` should return 13
- The array `[]` where n is `4` should return 0

## Rules
- If n is greater than the length of the array return 0

The following test cases will determine whether the above rules are applied:

```
		{"zero case", []int{}, 4, 0},
		{"windown length 2 array length 7", []int{1, 2, 5, 2, 8, 1, 5}, 2, 10},
		{"window length 4 array length 7", []int{1, 2, 5, 2, 8, 1, 5}, 4, 17},
		{"window length 1 array length 4", []int{4, 2, 1, 6}, 1, 6},
		{"window length 4 array length 5", []int{4, 2, 1, 6, 2}, 4, 13},
```

## Solution
- Check the length of the array, if it is less than n, return 0
- Use `maxSum` and `tempSum` to track the current max Sum and calculate the temp Sum
- Loop through the first window of `n` elements and set the value of `maxSum` to their sum
- Set `tempSum` to the value of `maxSum`
- Loop through the array starting at `n` through to the last element of the array
  - Subtract the value of the last index of the window from `tempSum` and add the value of the current index to `tempSum`
  - Set `maxSum` to the max value between `tempSum` and `maxSum`
- return `maxSum`
```go
func maxSubarraySum(xi []int, n int) int {
	if len(xi) < n {
		return 0
	}

	maxSum := 0

	for i := range n {
		maxSum += xi[i]
	}

    tempSum := maxSum

	for i := n; i < len(xi); i++ {
		tempSum = tempSum - xi[i-n] + xi[i]
		maxSum = max(maxSum, tempSum)
	}

	return maxSum
}
```
