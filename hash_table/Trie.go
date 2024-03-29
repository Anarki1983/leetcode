package leetcode

//Runtime: 120 ms, faster than 32.99% of Go online submissions for Implement Trie (Prefix Tree).
//Memory Usage: 15.9 MB, less than 64.46% of Go online submissions for Implement Trie (Prefix Tree).

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
type Trie struct {
	runeMap map[rune]*Trie
	isTail  bool
}

func Constructor() Trie {
	return Trie{runeMap: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	currentNode := this
	for _, r := range word {
		if _, ok := currentNode.runeMap[r]; !ok {
			t := Constructor()
			currentNode.runeMap[r] = &t
		}

		currentNode = currentNode.runeMap[r]
	}

	currentNode.isTail = true
}

func (this *Trie) Search(word string) bool {
	currentNode := this
	for _, r := range word {
		if _, ok := currentNode.runeMap[r]; !ok {
			return false
		}

		currentNode = currentNode.runeMap[r]
	}

	if !currentNode.isTail {
		return false
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this
	for _, r := range prefix {
		if _, ok := currentNode.runeMap[r]; !ok {
			return false
		}

		currentNode = currentNode.runeMap[r]
	}

	return true
}
