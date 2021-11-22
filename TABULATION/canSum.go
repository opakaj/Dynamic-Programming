package main

func gridTraveller(targetSum int, numbers []int) bool {
	table := make([]bool, targetSum+1)
	table[0] = true

	for i := 0; i <= len(table); i++ {
		if table[i] == true {
			for _, x := range numbers {
				table[i+x] = true
			}
		}
	}
	return table[targetSum]
}

//Time O(mn)
//Space O(m)
