package linkedlist

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DLL struct {
	Head *Node
}

func (l *DLL) Insert(x int) *Node {
	if l.Head == nil {
		l.Head = &Node{Data: x}
		return l.Head
	}
	if l.Head.Data > x {
		node := &Node{Data: x, Next: l.Head}
		l.Head.Prev = node
		l.Head = node
		return node
	}
	for n := l.Head; n != nil; n = n.Next {
		if n.Next == nil {
			n.Next = &Node{Data: x, Prev: n}
			return n.Next
		}
		if n.Next.Data > x {
			node := &Node{Data: x, Prev: n, Next: n.Next}
			n.Next.Prev = node
			n.Next = node
			return node
		}
	}
	return nil
}

func (l *DLL) Remove(x int) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == x {
		l.Head.Next.Prev = nil
		l.Head = l.Head.Next
	}
	for n := l.Head; n != nil; n = n.Next {
		if n.Data == x {
			if n.Next == nil {
				n.Prev.Next = nil
				return
			}
			n.Prev.Next = n.Next
			n.Next.Prev = n.Prev
			return
		}
	}
}

func (l *DLL) Print() {
	fmt.Print("HEAD -> ")
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%v <=> ", n.Data)
	}
	fmt.Print("\n")
}
