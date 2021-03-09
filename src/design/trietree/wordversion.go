package trietree

type WordTrieNode struct {
	value    byte
	children []*WordTrieNode
	isEnd    bool
}

type WordTrie struct {
	Root *WordTrieNode
}

/** Initialize your data structure here. */
func Constructor() WordTrie {
	t := WordTrie{}
	trieNode := new(WordTrieNode)
	trieNode.children = make([]*WordTrieNode, 26)
	t.Root = trieNode
	return t
}

/** Inserts a word into the trie. */
func (this *WordTrie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := this.Root
	for i := 0; i < len(word); i++ {
		key := word[i]
		if node.children[key-'a'] == nil {
			trieNode := new(WordTrieNode)
			trieNode.value = key
			trieNode.children = make([]*WordTrieNode, 26)
			node.children[key-'a'] = trieNode
		}
		node = node.children[key-'a']
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *WordTrie) Search(word string) bool {
	node := this.Root
	for i := 0; i < len(word); i++ {
		key := word[i]
		if node.children[key-'a'] == nil {
			return false
		}
		node = node.children[key-'a']
	}
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *WordTrie) StartsWith(prefix string) bool {
	node := this.Root
	for i := 0; i < len(prefix); i++ {
		key := prefix[i]
		if node.children[key-'a'] == nil {
			return false
		}
		node = node.children[key-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
