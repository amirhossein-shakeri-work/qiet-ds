package queue

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Q struct {
	Head *Node
}

func (q *Q) EnQ(x int) {
	if q.Head == nil {
		q.Head = &Node{Data: x}
		return
	}

	if q.Head.Data > x {
		q.Head = &Node{Data: x, Next: q.Head}
		return
	}
	for n := q.Head; n != nil; n = n.Next {
		if n.Next == nil {
			n.Next = &Node{Data: x}
			break
		}
		if n.Next.Data > x {
			n.Next = &Node{Data: x, Next: n.Next}
			break
		}
	}
}

func (q *Q) DeQ() *Node {
	if q.Head == nil {
		return nil
	}
	el := *q.Head
	q.Head = q.Head.Next
	return &el
}

func (q *Q) PrintAscending() string {
	str := "HEAD -> "
	for n := q.Head; n != nil; n = n.Next {
		str += fmt.Sprintf("%v -> ", n.Data)
	}
	return str
}
