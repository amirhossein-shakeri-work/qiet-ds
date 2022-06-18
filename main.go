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

	t0 := tree.NewExpTree(nil)
	t0.ParsePrefix("*+23-41")
	fmt.Println(t0.ToPrefix())
}
