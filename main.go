package main

import (
	"amirhossein-shakeri/qiet-ds/tree"
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

	t1 := tree.NewExpTree(nil)
	t1.ParseInfix("((5+7)/(9-6))")
	fmt.Println(t1.ToPostfix())
	fmt.Println(t1.Evaluate())
}
