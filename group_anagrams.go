package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	words := make(map[string][]string)
	var groupedAnagrams [][]string

	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(a, b int) bool {
			return runes[a] < runes[b]
		})
		sorted := string(runes)

		words[sorted] = append(words[sorted], word)
	}

	for _, value := range words {
		groupedAnagrams = append(groupedAnagrams, value)
	}

	return groupedAnagrams
}
