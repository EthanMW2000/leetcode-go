package main

func groupAnagrams(strs []string) [][]string {
	newArray := [][]string{}
	wordMap := make(map[byte]int)

	for i := 0; i < len(strs[0]); i++ {
		wordMap[strs[0][i]] = 0
	}

	newArray = append(newArray, []string{strs[0]})

	for len(strs) > 1 {

	}

	return [][]string{}
}

func sortAnagrams(strs []string, wordMap map[byte]int, newArray [][]string)