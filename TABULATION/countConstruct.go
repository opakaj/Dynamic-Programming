package main

func canConstruct(target string, wordBank []string) int {
	table := make([]int, len(target)+1)
	table[0] = 1 // with an empty string, there is only one way to generate the empty string

	for i := 0; i <= len(target); i++ {
		for _, word := range wordBank {
			if target[i:i+len(word)] == word {
				table[i+len(word)] += table[i]
			}
		}
	}
	return table[len(target)]
}

//Time O(m^2n)
//Space O(m)
