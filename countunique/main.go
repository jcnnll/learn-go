package main

import "fmt"

func main() {
	xi := []int{}
	fmt.Println(countUnique(xi))
}

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
