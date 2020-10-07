package remark

import "github.com/gomarkdown/markdown/ast"

type ListItem struct {
	*ast.ListItem
	parent Element
}

func (l *ListItem) Parent() Element {
	return l.parent
}

func (li ListItem) String() string {
	return ""
}

func NewListItem(node *ast.ListItem) (*ListItem, error) {
	li := &ListItem{
		ListItem: node,
	}
	return li, nil
}
