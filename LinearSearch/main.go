package main

import "fmt"

type IndVal struct {
	Index int
	Value int
}

func linearSearch(arr []int, tElemt int) IndVal {
	var result IndVal
	for i, v := range arr {
		if v == tElemt {
			result = IndVal{Index: i, Value: v}
		}
	}
	return result
}

func main() {
	arr := []int{5, 6, 4, 8, 7, 9, 5, 4, 6, 83}
	targetElement := 3
	fmt.Printf("%+v\n", linearSearch(arr, targetElement))
}
