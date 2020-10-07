package remark

import "github.com/gomarkdown/markdown/ast"

type List struct {
	*ast.List
	parent Element
}

func (l *List) Parent() Element {
	return l.parent
}

func (l List) String() string {
	return ""
}

func NewList(node *ast.List) (*List, error) {
	l := &List{
		List: node,
	}
	return l, nil
}
