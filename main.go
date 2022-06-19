package main

import (
	"amirhossein-shakeri/qiet-ds/linkedlist"
	"fmt"
)

func main() {
	fmt.Println("Starting")
	// t := tree.NewExpTree(nil)
	// t.ParseInfix("((2+3)*(4-1))")
	// fmt.Println(t.ToPrefix())

	// t0 := tree.NewExpTree(nil)
	// t0.ParsePrefix("*+23-41")
	// fmt.Println(t0.ToPrefix())
	// fmt.Println(t0.ToPostfix())

	// t1 := tree.NewExpTree(nil)
	// t1.ParseInfix("((5+7)/(9-6))")
	// fmt.Println(t1.ToPostfix())
	// fmt.Println(t1.Evaluate())

	// q1 := &queue.Q{}
	// q1.EnQ(4)
	// q1.EnQ(6)
	// q1.EnQ(5)
	// q1.EnQ(2)
	// q1.EnQ(1)
	// q1.EnQ(9)
	// q1.EnQ(3)
	// fmt.Println(q1.PrintAscending())
	// fmt.Printf("Deq: %v\n", q1.DeQ().Data)
	// fmt.Println(q1.PrintAscending())
	// fmt.Printf("Deq: %v\n", q1.DeQ().Data)
	// fmt.Println(q1.PrintAscending())

	l := &linkedlist.DLL{}
	l.Insert(3)
	l.Print()
	l.Insert(4)
	l.Print()
	l.Insert(2)
	l.Print()
	l.Insert(8)
	l.Insert(9)
	l.Print()
	l.Remove(4)
	l.Print()
	l.Remove(8)
	l.Print()
	l.Remove(2)
	l.Print()
	l.Insert(4)
	l.Print()
}
