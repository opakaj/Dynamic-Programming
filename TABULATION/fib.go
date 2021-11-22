package main

func fib(n int) int {
	table := make([]int, n+1)
	table[1] = 1

	for i := 0; i <= n; i++ {
		table[i+1] += table[i]
		table[i+2] += table[i]
	}
	return table[n]
}

//Time O(n)
//Space O(n)
