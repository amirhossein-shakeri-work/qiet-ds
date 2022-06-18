package main

import (
	"amirhossein-shakeri/qiet-ds/exptree"
	"fmt"
)

func main() {
	fmt.Println("Starting")
	t := exptree.New(nil)
	// t.ParseInfix("(2+3)")
	// t.Root.Print()
	t.ParseInfix("((2+3)*(4-1))")
	t.Root.Print()
}
