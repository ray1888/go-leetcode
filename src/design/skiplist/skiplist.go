package skiplist

import (
	"math"
	"math/rand"
)

const (
	maxLevel = 16
	maxRand  = 65535.0
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

type skipNode struct {
	value int
	right *skipNode
	down  *skipNode
}

type Skiplist struct {
	head *skipNode
}

func Constructor() Skiplist {
	// 初始化所有的层级的链表，最大跳表层级为16， 步长为2^n (0<=n<= 16)
	left := make([]*skipNode, maxLevel)
	right := make([]*skipNode, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &skipNode{-1, nil, nil}
		right[i] = &skipNode{20001, nil, nil}
	}
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		// 绑定与下一个层级的关系
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxLevel-1].right = right[maxLevel-1]
	return Skiplist{left[0]}
}

/*
   本质上是通过多级别的二分查找来减低查询的时间复杂度
*/
func (this *Skiplist) Search(target int) bool {
	node := this.head
	for node != nil {
		// 当比第一个小的时候，则走下一级别的链表，否则就继续往存在的链表下面走
		if node.right.value > target {
			node = node.down
		} else if node.right.value < target {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	prev := make([]*skipNode, maxLevel)
	i := 0
	node := this.head
	for node != nil {
		if node.right.value >= num {
			prev[i] = node
			i++
			node = node.down
		} else {
			node = node.right
		}
	}
	// 随机的去生成 Level，这个是 skipList 最重要的概念
	n := randLevel()
	arr := make([]*skipNode, n)
	t := &skipNode{-1, nil, nil}
	for i, a := range arr {
		p := prev[maxLevel-n+i]
		a = &skipNode{num, p.right, nil}
		p.right = a
		t.down = a
		t = a
	}
}

func (this *Skiplist) Erase(num int) bool {
	/*
	   找到对应的层级，本质上就是去把这个节点存在在每个层级的对应链表进行修改,移除相关的节点
	   从顶到底进行搜索，所以本质上维护这个跳表的各个层级平摊到每个插入和删除的操作中，
	   而不用单独每次遍历最底层的链表来重新生成，减低了维护的成本
	*/
	var ans bool
	node := this.head
	for node != nil {
		if node.right.value > num {
			node = node.down
		} else if node.right.value < num {
			node = node.right
		} else {
			ans = true
			node.right = node.right.right
			node = node.down
		}
	}
	return ans
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
