package main

import (
	"amirhossein-shakeri/qiet-ds/tree"
	"fmt"
)

func main() {
	fmt.Println("Starting")
	t := tree.NewExpTree(nil)
	// t.ParseInfix("(2+3)")
	// t.Root.Print()
	t.ParseInfix("((2+3)*(4-1))")
	fmt.Println(t.ToPrefix())
	t.Root.Print()
}
