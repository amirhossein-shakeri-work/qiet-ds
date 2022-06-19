package tree

import (
	"amirhossein-shakeri/qiet-ds/stack"
	"fmt"
)

type ExpTree struct {
	Root *Node
}

func NewExpTree(root *Node) *ExpTree {
	return &ExpTree{Root: root}
}

func (tree *ExpTree) ParseInfix(s string) *ExpTree { // "((2+3)*(4-1))"
	var n *Node
	root := Node{Data: "ROOT"}
	for i := 0; i < len(s)-1; i++ {
		ch := string(s[i])
		if ch == "(" {
			if n == nil {
				n = &root
			}
			n.LChild = &Node{Parent: n}
			n.RChild = &Node{Parent: n}
			n = n.LChild
			continue
		}
		if ch == ")" {
			n = n.Parent
			continue
		}
		if ch == " " {
			continue
		}
		if isOperator(ch) {
			n.Data = ch
			n = n.RChild
		} else {
			n.Data = ch
			n = n.Parent
		}
	}
	tree.Root = n
	return tree
}

func (tree *ExpTree) ParsePrefix(s string) *ExpTree { // "*+23-41"
	i := 0
	tree.Root = PrefixToExpTree(s, &i)
	return tree
}

func (tree *ExpTree) ToPrefix() string {
	return *tree.Root.ToPrefix()
}

func (tree *ExpTree) ToPostfix() string {
	return *tree.Root.ToPostfix()
}

func (tree *ExpTree) Evaluate() int {
	str := tree.ToPostfix()
	stack := &stack.Stack{}
	var value int
	for _, ch := range str {
		if isOperator(string(ch)) {
			n2 := stack.Pop()
			n1 := stack.Pop()
			switch string(ch) {
			case "+":
				value = n1 + n2
			case "-":
				value = n1 - n2
			case "*":
				value = n1 * n2
			case "/":
				value = n1 / n2
			}
			stack.Push(value)
		} else {
			stack.Push(int(ch - '0'))
		}
	}
	return value
}

func (t *ExpTree) Print() {
	fmt.Printf("%#v", t)
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}
