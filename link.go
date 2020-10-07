package remark

import "github.com/gomarkdown/markdown/ast"

type Link struct {
	*ast.Link
	parent Element
}

func (l *Link) Parent() Element {
	return l.parent
}

func (l Link) String() string {
	return ""
}

func NewLink(node *ast.Link) (*Link, error) {
	l := &Link{
		Link: node,
	}
	return l, nil
}
