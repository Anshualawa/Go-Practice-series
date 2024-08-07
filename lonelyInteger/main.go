package main

import (
	"fmt"
)

func findUnique(elements []int) []int {
	counts := make(map[int]int)
	uniq := []int{}
	// Count occurrences of each element
	for _, elem := range elements {
		counts[elem]++
	}
	fmt.Println(counts)
	// Find the element that occurs only once
	for elem, count := range counts {
		if count == 1 {
			uniq = append(uniq, elem)
		}
	}

	return uniq
}

func main() {
	a := []int{1, 2, 8, 3, 4, 3, 2, 8, 1, 0, 8}
	unique := findUnique(a)
	fmt.Printf("The unique element is: %d\n", unique)
}
