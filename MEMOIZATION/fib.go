package main

import "fmt"

func fib(n int, memo map[int]int) int {
	//if the n is in the memeo
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func main() {
	m := map[int]int{}
	fmt.Println(fib(6, m))
	fmt.Println(fib(7, m))
	fmt.Println(fib(8, m))
	fmt.Println(fib(50, m))
}

//time O(2^n)
//space O(n)
