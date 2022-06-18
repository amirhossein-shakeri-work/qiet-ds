package tree

import "fmt"

type Node struct {
	Data   string
	LChild *Node
	RChild *Node
	Parent *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Print() {
	fmt.Printf("%#v", n)
}

func (n *Node) ToPrefix() *string {
	if n == nil {
		return nil
	}
	s := n.Data
	if n.LChild != nil {
		s += *n.LChild.ToPrefix()
	}
	if n.RChild != nil {
		s += *n.RChild.ToPrefix()
	}
	return &s
}

func PrefixToExpTree(s string, i *int) *Node {
	if *i >= len(s) {
		return nil
	}
	n := &Node{Data: string(s[*i])}
	if isOperator(string(s[*i])) {
		*i++
		n.LChild = PrefixToExpTree(s, i)
		*i++
		n.RChild = PrefixToExpTree(s, i)
	}
	return n
}
