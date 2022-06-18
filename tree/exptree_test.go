package tree_test

import (
	"amirhossein-shakeri/qiet-ds/tree"
	"testing"
)

func TestParseInfix(t *testing.T) {
	infixExp := "((2+3) * (4-1))"
	expected := &tree.ExpTree{
		Root: &tree.Node{
			Data: "*",
			LChild: &tree.Node{
				Data:   "+",
				LChild: &tree.Node{Data: "2"},
				RChild: &tree.Node{Data: "3"},
			},
			RChild: &tree.Node{
				Data:   "-",
				LChild: &tree.Node{Data: "4"},
				RChild: &tree.Node{Data: "1"},
			},
		},
	}
	tree := tree.NewExpTree(nil)
	tree.ParseInfix(infixExp)
	if tree.Root.Data != "*" || tree.Root.RChild.Data != "-" {
		t.Errorf("Trees don't match. Expected: %v", expected)
	}
}
