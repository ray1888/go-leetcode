package lru

import (
	"testing"
)

const (
	LRUlength  = 4
	OverLength = 5
)

func TestLRUGetUnexisted(t *testing.T) {
	lru, _ := NewLRUCache(LRUlength)
	lru.Put(4, 4)
	val := lru.Get(3)
	if val != -1 {
		t.Fatal("get unexisted failed, not matched")
	}
}

func TestLRUSaveGet(t *testing.T) {
	lru, _ := NewLRUCache(LRUlength)
	lru.Put(4, 4)
	val := lru.Get(4)
	if val != 4 {
		t.Fatal("get puted val failed, not matched")
	}
}

func TestLRUSaveGetwithLength(t *testing.T) {
	lru, _ := NewLRUCache(LRUlength)
	for i := 0; i < LRUlength; i++ {
		lru.Put(i, i)
	}
	for i := 0; i < LRUlength; i++ {
		val := lru.Get(i)
		if val != i {
			t.Fatal("lru can't get puted val within length")
		}
	}
}

func TestLRUSaveGetOverLength(t *testing.T) {
	lru, _ := NewLRUCache(LRUlength)
	for i := 0; i < OverLength; i++ {
		lru.Put(i, i)
	}
	for i := 0; i < OverLength; i++ {
		val := lru.Get(i)
		if i == 0 && val == -1 {
			continue
		}
		if val != i {
			t.Fatalf("lru can't get puted val within length, key is %d, value is %d", i, val)
		}
	}
}
