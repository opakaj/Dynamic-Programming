package main

func gridTraveller(m, n int) int {
	var table [][]int
	table[1][1] = 1

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			current := table[0][0]
			if j+1 <= n {
				table[i][j+1] += current
			}
			if i+1 <= m {
				table[i+1][j] += current
			}
		}
	}
	return table[m][n]
}

//Time O(mn)
//Space O(mn)
