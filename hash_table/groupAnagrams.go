package hash_table

import "sort"

// string length: m
// Time complexity: O(n*mlogm)
// Space complexity: O(n)

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0, 2)
	anagramMap := make(map[string][]string)

	// sort rune as key of anagramMap, then append origin word to value slice
	for i := 0; i < len(strs); i++ {
		key := genSortedStr(strs[i])

		anagramMap[key] = append(anagramMap[key], strs[i])
	}

	for _, anagram := range anagramMap {
		ret = append(ret, anagram)
	}

	return ret
}

func genSortedStr(word string) string {
	runes := []rune(word)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
