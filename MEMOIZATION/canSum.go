package main

import "fmt"

func canSum(targetSum int, numbers []int, memo map[int]bool) bool {
	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return true
	}
	//cater for negative remainders/target sums
	if targetSum < 0 {
		return false
	}

	for _, num := range numbers {
		remainder := targetSum - num
		if canSum(remainder, numbers, memo) == true {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}

func main() {
	memo := map[int]bool{}
	fmt.Println(canSum(7, []int{2, 3}, memo))
	fmt.Println(canSum(7, []int{5, 4, 3, 7}, memo))
	fmt.Println(canSum(7, []int{2, 4}, memo))
	fmt.Println(canSum(8, []int{2, 3, 5}, memo))
	fmt.Println(canSum(300, []int{7, 14}, memo))
}
