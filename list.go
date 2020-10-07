package remark

import "github.com/gomarkdown/markdown/ast"

type List struct {
	*ast.List
}

func (l List) String() string {
	return ""
}

func NewList(node *ast.List) (List, error) {
	l := List{
		List: node,
	}
	return l, nil
}
