package lfu

type LFUCache struct {
	cache               map[int]*Node
	freq                map[int]*DoubleList
	ncap, size, minFreq int
}

/*
	本质上，要解决LFU的问题，首先要定义问题，
	因为题目给到本质上他是由两个维度去组成来做缓存的更新
	1. 使用频次 ， 2. 最近使用的时间的长短
	因此这里就需要有两个维度的数据结构来存储这种数据
	那解法有有两种了：
	1. 一个哈希表进行cache查找，并且使用堆来维护数据的顺序，时间复杂度为O(logN)
	2. 使用两个哈希表来进行查找：
		一个是用于缓存使用过的频次和节点顺序的关系
		另一个是用于key和实际存储节点的映射
		这样就能维护两个维度的数据
	然后实现上使用了双向链表来维护存储频次和节点顺序的关系的原因是，
	对于双向链表，我们可以通过插入的顺序来表达他实际的顺序，而不用额外维护更多的索引的信息
	因此：
	每次发生使用频次变动的操作，本质上是
	1. 从原来的频次链表中删除这个节点
	2. 修改节点的使用频次，并且检查新生成的频次是否已经存在，然后插入头部
*/

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache: make(map[int]*Node),
		freq:  make(map[int]*DoubleList),
		ncap:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.IncFreq(node)
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key, value int) {
	if this.ncap == 0 {
		return
	}
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.IncFreq(node)
	} else {
		if this.size >= this.ncap {
			node := this.freq[this.minFreq].RemoveLast()
			delete(this.cache, node.key)
			this.size--
		}
		x := &Node{key: key, val: value, freq: 1}
		this.cache[key] = x
		if this.freq[1] == nil {
			this.freq[1] = CreateDL()
		}
		this.freq[1].AddFirst(x)
		this.minFreq = 1
		this.size++
	}
}

func (this *LFUCache) IncFreq(node *Node) {
	_freq := node.freq
	this.freq[_freq].Remove(node)
	/*
	 此部分如果不是为了节省内存可以不使用，
	 因为每次创建和删除这样的key和双向链表的Map每次都要单独的allcoate内存，对性能伤害大
	*/
	if this.minFreq == _freq && this.freq[_freq].IsEmpty() {
		this.minFreq++
		delete(this.freq, _freq)
	}

	node.freq++
	if this.freq[node.freq] == nil {
		this.freq[node.freq] = CreateDL()
	}
	this.freq[node.freq].AddFirst(node)
}

type DoubleList struct {
	head, tail *Node
}

type Node struct {
	prev, next     *Node
	key, val, freq int
}

func CreateDL() *DoubleList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return &DoubleList{
		head: head,
		tail: tail,
	}
}

func (this *DoubleList) AddFirst(node *Node) {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

func (this *DoubleList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil
}

func (this *DoubleList) RemoveLast() *Node {
	if this.IsEmpty() {
		return nil
	}

	last := this.tail.prev
	this.Remove(last)

	return last
}

func (this *DoubleList) IsEmpty() bool {
	return this.head.next == this.tail
}
