package mapRelated

import (
	"math/rand"
)

type DataItem struct {
	Val   int
	Index int
}

type RandomizedCollection2 struct {
	data     []DataItem
	indexMap map[int][]int
}

/** Initialize your data structure here. */
func Constructor2() RandomizedCollection2 {
	item := RandomizedCollection2{}
	item.data = make([]DataItem, 0)
	item.indexMap = make(map[int][]int)
	return item
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection2) Insert(val int) bool {
	exist := true
	if _, ok := this.indexMap[val]; ok {
		exist = false
	} else {
		this.indexMap[val] = make([]int, 0)
	}
	this.data = append(this.data, DataItem{val, len(this.indexMap[val])})
	this.indexMap[val] = append(this.indexMap[val], len(this.data)-1)
	return exist
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection2) Remove(val int) bool {
	if _, ok := this.indexMap[val]; !ok {
		return false
	}
	lastIdx := len(this.data) - 1
	idx := this.indexMap[val][len(this.indexMap[val])-1]
	this.indexMap[val] = this.indexMap[val][:len(this.indexMap[val])-1]
	if lastIdx != idx {
		lastPair := this.data[lastIdx]
		lastVal := lastPair.Val
		lastValPosInMap := lastPair.Index
		this.data[idx] = lastPair
		this.indexMap[lastVal][lastValPosInMap] = idx
	}
	this.data = this.data[:lastIdx]
	if len(this.indexMap[val]) == 0 {
		delete(this.indexMap, val)
	}
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection2) GetRandom() int {
	if len(this.data) > 0 {
		r := rand.Intn(len(this.data))
		return this.data[r].Val
	}
	return -1
}
