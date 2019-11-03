package main

import "fmt"

// Time: O(n)
// Space: O(n)
//
// This is basic fibonacci plus memoization to make the algorithm
// linear time.
func fibonacci(n int) int {
	return doFibonacci(n, map[int]int{})
}

func doFibonacci(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = doFibonacci(n-1, memo) + doFibonacci(n-2, memo)
	}
	return memo[n]
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
