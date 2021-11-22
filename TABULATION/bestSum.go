package main

func bestSum(targetSum int, numbers []int) []int {
	table := make(map[int][]int, targetSum+1)
	table[0] = make([]int, 1)

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, num := range numbers {
				combination := append(table[i], num)
				if table[i] != nil || len(table[i]) > len(combination) {
					table[i+num] = combination
				}
			}
		}
	}
	return table[targetSum]
}

//Time O(m^2n)
//Space O(m^2)
