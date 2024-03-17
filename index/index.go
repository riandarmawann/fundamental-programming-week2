package main

import "fmt"

func findIndex(x int, arr []int) []int {
	var indices []int

	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			indices = append(indices, i)
		}
	}

	return indices
}

func main() {
	fmt.Println(findIndex(5, []int{1, 5, 8, 9, 2, 3, 5, 9, 12}))
}
