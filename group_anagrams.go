package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	newArray := [][]string{}
	wordMaps := []map[byte]int{}
	seenIndexes := make(map[int]int)

	for i := 0; i < len(strs); i++ {
		wordMaps = append(wordMaps, map[byte]int{})
		for j := 0; j < len(strs[i]); j++ {
			wordMaps[i][strs[i][j]] = 1
		}
	}

	newArraySpot := 0
	current := 0
	for len(seenIndexes) != len(strs) {
		fmt.Printf("seenIndexes %d, strs %d", len(seenIndexes), len(strs))
		newArray = append(newArray, []string{})
		newArray[newArraySpot] = append(newArray[newArraySpot], strs[current])
		seenIndexes[current] = 1

		for i := current + 1; i < len(strs); i++ {
			if seenIndexes[i] == 1 {
				continue
			}

			innerCount := 1
			for key := range wordMaps[i] {
				if _, ok := wordMaps[current][key]; !ok {
					break
				}

				if innerCount == len(wordMaps[i]) {
					newArray[newArraySpot] = append(newArray[newArraySpot], strs[i])
					seenIndexes[i] = 1
				}
				innerCount++
			}
		}

		for seenIndexes[current] == 1 {
			current++
		}
		newArraySpot++
	}

	return newArray
}

func main() {
	anagrams := []string{"eat", "ate", "bat", "nat", "tea", "tan"}
	groupedAnagrams := groupAnagrams(anagrams)

	fmt.Println(groupedAnagrams)
}
