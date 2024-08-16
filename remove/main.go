package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "bolvvvv" // bovvvlet    bovvvvvlet   bovvvvwwlet   bolvvvv
	fmt.Println(removeDuplicateChar(input))
	// expected  output "bolet"
}

func removeDuplicateChar(s string) string {
	str := strings.Split(s, "")
	seen := make(map[string]bool)
	var result []string

	for _, char := range str {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}
	return strings.Join(result, "")
}
