package main

import (
	"fmt"
	"sort"
)

func binarySerach(arr []int, targetVal int) bool {
	found := false
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == targetVal {
			found = true
			break
		}
		if arr[mid] > targetVal {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return found
}
func main() {
	var arr = []int{1, 8, 2, 3, 4, 5, 6, 7, 9}
	sort.Ints(arr)
	fmt.Println(binarySerach(arr, 8))
}
