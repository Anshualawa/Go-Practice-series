package main

import "fmt"

// Checks if a number is prime
func isPrime(num int) bool {
	if num <= 1 { // Numbers less than or equal to 1 are not prime
		return false
	}
	for i := 2; i*i <= num; i++ { // Check divisibility up to the square root of num
		if num%i == 0 { // If divisible by any number other than 1 and itself, not prime
			return false
		}
	}
	return true // If no divisors found, the number is prime
}

func main() {
	var container = []int{}   // Initialize an empty slice to store prime numbers
	i := 0                    // Start checking numbers from 0
	for len(container) <= 6 { // Continue until we have 6 prime numbers
		if isPrime(i) { // If the current number is prime
			container = append(container, i) // Add it to the container
		}
		i++ // Move to the next number
	}
	fmt.Println(container) // Print the list of prime numbers
}
