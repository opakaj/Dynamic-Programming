package main

func canConstruct(target string, wordBank []string) bool {
	table := make([]bool, len(target)+1)
	table[0] = true

	for i := 0; i <= len(target); i++ {
		if table[i] == true {
			for _, word := range wordBank {
				//if the word matches the xter starting at position i
				if target[i:i+len(word)] == word {
					table[i+len(word)] = true
				}
			}
		}
	}
	return table[len(target)]
}

//Time O(m^2n)
//Space O(m)
