package main

import "fmt"

func fibonacci_tabulation(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	lenth := n + 1
	var F = make([]int, lenth)
	F[0] = 0
	F[1] = 1
	for i := 2; i < lenth; i++ {
		F[i] = F[i-1] + F[i-2]
		// fmt.Println(F)
	}
	return F[n-1]
}

func main() {
	n := 15
	result := fibonacci_tabulation(n)
	fmt.Printf("The %vth Fibonacci number is %v\n", n, result)
}
