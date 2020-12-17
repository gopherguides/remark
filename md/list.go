package md

import (
	"fmt"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func NewList(node *ast.List) (*remark.List, error) {
	l := &remark.List{}

	switch node.ListFlags {
	case ast.ListTypeOrdered:
		l.Type = ListTypeOrdered
	case ast.ListTypeUnordered:
		l.Type = ListTypeUnordered
	default:
		return nil, fmt.Errorf("unknown list type %v", node)
	}

	for _, n := range node.Children {
		// always li
		el, err := NewElement(n)
		if err != nil {
			return nil, err
		}
		l.Elements = append(l.Elements, el)
	}
	return l, nil
}
