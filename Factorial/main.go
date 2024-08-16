package main

import "fmt"

func Factorial(num int64) int64 {
	if num == 0 {
		return 1
	}
	return num * Factorial(num-1)
}

func Fibonacci(num int64) []int64 {
	list := []int64{}
	list = append(list, 0, 1)

	for i := 2; list[len(list)-1] <= num; i++ {
		list = append(list, list[i-2]+list[i-1])
	}

	return list[:len(list)-1]
}

func main() {
	fmt.Println(Factorial(5))
	fmt.Println(Fibonacci(90))
}
