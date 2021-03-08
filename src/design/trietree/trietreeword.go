package trietree

import "fmt"

type runeTireNode struct {
	child map[rune]*runeTireNode
	Vaule interface{}
}
type RuneTire struct {
	root *runeTireNode
}

func NewRuneTire() *RuneTire {
	return &RuneTire{root: &runeTireNode{child: make(map[rune]*runeTireNode)}}
}
func (r *RuneTire) Insert(key string, value interface{}) {
	node := r.root
	for _, c := range key {
		if n, ok := node.child[c]; !ok {
			newNode := &runeTireNode{child: make(map[rune]*runeTireNode)}
			node.child[c] = newNode
			node = newNode
		} else {
			node = n
		}
	}
	node.Vaule = value
}

func (r *RuneTire) Get(key string) interface{} {
	node := r.root
	for _, c := range key {
		if n, ok := node.child[c]; !ok {
			return nil
		} else {
			node = n
		}
	}
	return node.Vaule
}
func main() {
	r := NewRuneTire()
	r.Insert("河北", "河北")
	r.Insert("湖南", "湖南")
	r.Insert("湖北", "湖北省")
	fmt.Println(r.Get("湖北"))
}
