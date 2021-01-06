package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewList(node *ast.List) (remark.Tag, error) {
	var t remark.Tag

	switch node.ListFlags {
	case ast.ListTypeOrdered:
		t = &remark.OrderedList{}
	default:
		t = &remark.OrderedList{}
	}

	err := set(t, node)
	return t, err
}
