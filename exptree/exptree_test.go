package exptree_test

import (
	"amirhossein-shakeri/qiet-ds/exptree"
	"testing"
)

func TestParseInfix(t *testing.T) {
	infixExp := "((2+3) * (4-1))"
	expected := &exptree.ExpTree{
		Root: &exptree.Node{
			Data: "*",
			LChild: &exptree.Node{
				Data:   "+",
				LChild: &exptree.Node{Data: "2"},
				RChild: &exptree.Node{Data: "3"},
			},
			RChild: &exptree.Node{
				Data:   "-",
				LChild: &exptree.Node{Data: "4"},
				RChild: &exptree.Node{Data: "1"},
			},
		},
	}
	tree := exptree.New(nil)
	tree.ParseInfix(infixExp)
	if tree.Root.Data != "*" || tree.Root.RChild.Data != "-" {
		t.Errorf("Trees don't match. Expected: %v", expected)
	}
}
