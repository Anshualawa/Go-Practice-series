package main

import "fmt"

func binarySerach(arr []int64, targetVal int64) bool {
	found := false
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		//fmt.Println(mid)
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
	var arr = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySerach(arr, 3))
}
