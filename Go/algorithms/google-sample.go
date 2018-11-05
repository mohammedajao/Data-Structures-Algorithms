package main

import "fmt"

func fibonacci_sequence(n int, memo map[int]int) int { // Recursive O(n) solution with memoization
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		memo[n] = fibonacci_sequence(n-1, memo) + fibonacci_sequence(n-2, memo)
	}
	fmt.Println(memo)
	return memo[n]
}

func main() {
	m := make(map[int]int)
	println(fibonacci_sequence(6, m))
}
