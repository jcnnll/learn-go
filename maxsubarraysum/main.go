package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySum([]int{}, 2))
	fmt.Println(maxSubarraySum([]int{1, 2, 5, 2, 8, 1, 5}, 2))
	fmt.Println(maxSubarraySum([]int{1, 2, 5, 2, 8, 1, 5}, 4))
}

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
