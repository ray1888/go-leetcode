package trietree

import (
	"fmt"
	"testing"
)

func TestTrieInsertSearch(t *testing.T) {
	trie := NewTrie()
	url := []string{"/main", "/process"}
	trie.Insert(url)
	result := trie.Search(url)
	fmt.Printf("result is %v\n", result)
	if !result {
		t.Fatal("can't match correct path \n")
	}
}

func TestTrieStartswith(t *testing.T) {
	trie := NewTrie()
	url := []string{"/main", "/process"}
	trie.Insert(url)
	surl := []string{"/main"}
	result := trie.Startswith(surl)
	fmt.Printf("result is %v", result)
	if !result {
		t.Fatal("can't match correct prefix path ")
	}
}
