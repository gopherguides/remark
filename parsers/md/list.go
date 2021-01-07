package md

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gopherguides/remark"
)

func (p *Parser) NewList(node *ast.List) (remark.Tag, error) {
	var t remark.Tag

	switch node.ListFlags {
	case ast.ListTypeOrdered:
		t = &remark.OrderedList{}
	default:
		t = &remark.UnorderedList{}
	}

	err := p.set(t, node)
	return t, err
}
