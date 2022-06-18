package exptree

import "fmt"

type Node struct {
	Data   string
	LChild *Node
	RChild *Node
	Parent *Node
}

type ExpTree struct {
	Root *Node
}

func New(root *Node) *ExpTree {
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

func (n *Node) Print() {
	fmt.Printf("%#v", n)
}

func (t *ExpTree) Print() {
	fmt.Printf("%#v", t)
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "^"
}
