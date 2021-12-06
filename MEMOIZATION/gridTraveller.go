package main

import "fmt"

func gridTraveller(m, n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	memo[n] = gridTraveller(m-1, n, memo) + gridTraveller(m, n-1, memo)
	return memo[n]
}

func main() {
	memo := map[int]int{}
	fmt.Println(gridTraveller(1, 1, memo))
	fmt.Println(gridTraveller(2, 3, memo))
	fmt.Println(gridTraveller(3, 2, memo))
	fmt.Println(gridTraveller(3, 3, memo))
	fmt.Println(gridTraveller(18, 18, memo))
}
