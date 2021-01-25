package hash

type MyHashMap struct {
	b      int
	bucket [][]kv
}

type kv struct {
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{b: 1 << 10, bucket: make([][]kv, 1<<10)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	// 此处选择取模的数为1024-1=1023 的原因是因为素数
	// 此文章有做解析 https://cloud.tencent.com/developer/ask/79043
	index := key & (this.b - 1)
	for e := range this.bucket[index] {
		if this.bucket[index][e].key == key {
			this.bucket[index][e].value = value
			return
		}
	}
	this.bucket[index] = append(this.bucket[index], kv{key: key, value: value})
	return
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := key & (this.b - 1)
	for e := range this.bucket[index] {
		if this.bucket[index][e].key == key {
			return this.bucket[index][e].value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := key & (this.b - 1)
	removeIndex := 0
	remove := false
	for e := range this.bucket[index] {
		if this.bucket[index][e].key == key {
			removeIndex = e
			remove = true
			break
		}
	}
	if remove {
		this.bucket[index] = append(this.bucket[index][:removeIndex], this.bucket[index][removeIndex+1:]...)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
