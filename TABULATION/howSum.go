package main

func howSum(targetSum int, numbers []int) []int {
	table := make(map[int][]int, targetSum+1)
	table[0] = make([]int, 1)

	for i := 0; i <= targetSum; i++ {
		if table[i] != nil {
			for _, num := range numbers {
				table[i+num] = append(table[i], num)
			}
		}
	}
	return table[targetSum]
}
