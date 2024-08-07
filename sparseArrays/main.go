package main

import (
	"fmt"
	"strings"
)

func matchingStrings(stringsList []string, queries []string) []int32 {
	result := make([]int32, len(queries))
	for i, query := range queries {
		count := int32(0)
		for _, str := range stringsList {
			// Trim spaces from both the string and the query
			if strings.TrimSpace(str) == query {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func main() {
	// Example input
	stringsList := []string{"ab", " ab", "abc"}
	queries := []string{"ab", "abc", " bc"}

	// Call the function
	results := matchingStrings(stringsList, queries)

	// Print the results
	fmt.Println(results) // Expected Output: [2, 1, 0]
}
