package stringProblem

import "strconv"

type operation struct {
	op       byte
	priority int
}

func calculate(s string) int {
	s += "+"
	//实现两个栈，一个放数字，一个放操作数
	num_stack := construct()
	op_stack := construct()
	//实现一个hash表，存放不同操作符对应的优先级
	hash := make(map[byte]int, 0)
	hash['+'] = 1
	hash['-'] = 1
	//定义优先级
	prio := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		//如果是操作符
		if _, ok := hash[s[i]]; ok {
			oper := operation{s[i], hash[s[i]] + prio}
			for !op_stack.empty() && op_stack.peek().(operation).priority >= oper.priority {
				op := op_stack.pop().(operation).op
				cur := 0
				b := num_stack.pop().(int)
				a := num_stack.pop().(int)
				switch op {
				case '+':
					cur = a + b
				case '-':
					cur = a - b
				}
				num_stack.push(cur)
			}
			op_stack.push(oper)
		} else if s[i] == '(' {
			prio += 10
		} else if s[i] == ')' {
			prio -= 10
		} else {
			//是数字(数字可能右多位数)
			tmp := string(s[i])
			for s[i+1] >= '0' && s[i+1] <= '9' {
				tmp += string(s[i+1])
				i++
			}
			num, _ := strconv.Atoi(tmp)
			num_stack.push(num)
		}
	}
	return num_stack.pop().(int)
}

//用链表实现栈
type Node struct {
	Val  interface{}
	Next *Node
}

type stack struct {
	dummy *Node
}

func construct() *stack {
	return &stack{&Node{}}
}

func (s *stack) push(x interface{}) {
	node := &Node{x, nil}
	node.Next = s.dummy.Next
	s.dummy.Next = node
}

func (s *stack) pop() interface{} {
	val := s.dummy.Next.Val
	s.dummy.Next = s.dummy.Next.Next
	return val
}

func (s *stack) peek() interface{} {
	return s.dummy.Next.Val
}

func (s *stack) empty() bool {
	return s.dummy.Next == nil
}
