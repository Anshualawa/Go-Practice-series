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
	var n int = 10
	/*result := fibonacci_tabulation(n)
	fmt.Printf("The %vth Fibonacci number is %v\n", n, result)
	*/

	result := fibonaccinum(n)
	fmt.Println(result)
}

func fibonaccinum(n int) []int {
	result := make([]int, 0)
	F := make([]int, 0)
	F = append(F, 0, 1)
	i := 2
	for F[i-1] <= n {
		F = append(F, F[i-1]+F[i-2])
		if F[i]%2 == 0 {
			result = append(result, F[i])
		}
		i++
	}
	rt := len(result)
	if n <= 10 {
		return result[:rt]
	}
	return result[:rt-1]
}
