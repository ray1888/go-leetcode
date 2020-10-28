package mapRelated

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	M map[int]int
	D []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	r := RandomizedSet{}
	r.M = make(map[int]int)
	r.D = make([]int, 0)
	return r
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.M[val]; ok {
		return false
	}
	index := len(this.D)
	this.D = append(this.D, val)
	this.M[val] = index
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.M[val]; !ok {
		return false
	}
	new := this.D[len(this.D)-1]
	index := this.M[val]
	this.D[index] = new
	this.D = this.D[:len(this.D)-1]
	this.M[new] = index
	delete(this.M, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.D) > 0 {
		r := rand.Intn(len(this.D))
		fmt.Println(r)
		return this.D[r]
	}
	return -1
}
