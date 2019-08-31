package trietree

import "fmt"

// This Trie tree is for url matching

type TrieNode struct {
	children map[string]*TrieNode
	val      string
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	root := TrieNode{}
	root.children = make(map[string]*TrieNode)
	trie := Trie{}
	trie.root = &root
	return &trie
}

func (t *Trie) Insert(url []string) {
	cur := t.root
	for i := 0; i < len(url); i++ {
		if node, ok := cur.children[url[i]]; ok {
			cur = node
			continue
		} else {
			newNode := TrieNode{val: url[i]}
			newNode.children = make(map[string]*TrieNode)
			cur.children[url[i]] = &newNode
			fmt.Printf("insert url subnode %s\n", url[i])
			cur = &newNode
		}
	}
	cur.isEnd = true
	fmt.Printf("insert end, set node %v isend True \n", cur)
}

func (t *Trie) Search(url []string) bool {
	node := t.SearchEndPoint(url)
	return node != nil && node.isEnd == true
}

func (t *Trie) SearchEndPoint(url []string) *TrieNode {
	cur := t.root
	for i := 0; i < len(url); i++ {
		if node, ok := cur.children[url[i]]; ok {
			cur = node
		} else {
			cur = node
			fmt.Printf("can't find node, return nil\n")
			return cur
		}
	}
	fmt.Printf("find url %v\n", url)
	fmt.Printf("node is %v\n", cur)
	return cur
}

func (t *Trie) Startswith(url []string) bool {
	node := t.SearchEndPoint(url)
	return node != nil
}

//func(t *Trie)Delete(url []string) bool{
//
//}
